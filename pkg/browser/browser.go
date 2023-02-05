package browser

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/google/uuid"
)

type Browser struct {
	browser *rod.Browser
	data    *data
	Wg      *sync.WaitGroup
	Q       chan struct{}
}

func NewDataStore() *data {
	d := new(data)
	d.store = make(map[string]string)
	return d
}

type data struct {
	store map[string]string
	sync.RWMutex
}

func (d *data) Get(k string) (v string, ok bool) {
	d.RWMutex.RLock()
	defer d.RWMutex.RUnlock()
	v, ok = d.store[k]
	return v, ok
}

func (d *data) Add(k, v string) {
	d.RWMutex.Lock()
	defer d.RWMutex.Unlock()
	d.store[k] = v
}

func NewBrowser(flagDisableHeadless bool) (*Browser, error) {
	flagHeadless := !flagDisableHeadless
	d := NewIPadPro("en-US;q=0.9")
	u, err := launcher.New().
		Set("no-sandbox", "true").
		Headless(flagHeadless).
		Launch()
	if err != nil {
		return nil, err
	}

	browser := rod.New().ControlURL(u).Trace(true).DefaultDevice(d)
	err = browser.Connect()
	if err != nil {
		return nil, err
	}

	return &Browser{
		browser, NewDataStore(), &sync.WaitGroup{}, make(chan struct{}, 9),
	}, nil
}

func (b *Browser) ExtractVirtualDOM(url string) (err error) {
	page, err := b.browser.Page(proto.TargetCreateTarget{})
	if err != nil {
		fmt.Printf("Browser.Page()=%v\n", err)
		return err
	}

	err = page.Timeout(30 * time.Second).Navigate(url)
	if err != nil {
		fmt.Printf("page.Navigate()=%v\n", err)
		return err
	}
	if err := page.WaitLoad(); err != nil {
		fmt.Printf("page.WaitLoad()= %v\n", err)
		return err
	}

	htmlVirtualDOM, err := page.Element("html")
	if err != nil {
		fmt.Printf("page.Element()=%v\n", err)
		return err
	}

	domData, err := htmlVirtualDOM.HTML()
	if err != nil {
		fmt.Printf("htmlVirtualDOM.HTML()=%v\n", err)
		return err
	}

	fmt.Printf("%s\n", domData)
	return nil
}

func (b *Browser) ExtractAllResources(dirPath, url string) (err error) {
	fmt.Printf("[start] %s\n", url)

	page, err := b.browser.Page(proto.TargetCreateTarget{})
	if err != nil {
		fmt.Printf("Browser.Page()=%v\n", err)
		return err
	}

	router := b.browser.HijackRequests()
	router.MustAdd("*", b.HIjack(dirPath))
	go router.Run()
	defer func() {
		page.Close()
	}()
	defer func() {
		if err := router.Stop(); err != nil {
			fmt.Printf("router.Stop()=%v\n", err)
		}
	}()

	err = page.Timeout(30 * time.Second).Navigate(url)
	if err != nil {
		fmt.Printf("page.Navigate()=%v\n", err)
		return err
	}
	if err := page.WaitLoad(); err != nil {
		fmt.Printf("page.WaitLoad()= %v\n", err)
		return err
	}

	htmlVirtualDOM, err := page.Element("html")
	if err != nil {
		fmt.Printf("page.Element()=%v\n", err)
		return err
	}

	domData, err := htmlVirtualDOM.HTML()
	if err != nil {
		fmt.Printf("htmlVirtualDOM.HTML()=%v\n", err)
		return err
	}
	utils.Sleep(1)
	localStorage, err := page.Eval(`()=>{
		var result=""
		for(var i=0;i< localStorage.length;i++){
			var key=localStorage.key(i);
			var value=localStorage.getItem(key);
			result += "local_storage: "+key + " -> " + value + "\n";
		}
		return result;
	}`)
	if err != nil {
		fmt.Printf("Try to get localstorage instance: %v\n", err)
		return err
	}

	sessionStorage, err := page.Eval(`()=>{
		var result=""
		for (const [key, value] of Object.entries(sessionStorage)) {
			result += "session_storage: " + key + " -> " + value + "\n";
		};
		return result;
	}`)
	if err != nil {
		fmt.Printf("Try to get localstorage instance: %v\n", err)
		return err
	}

	fileName, ok := b.data.Get(url)
	if ok {
		path := filepath.Join(dirPath, fileName)

		f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Printf("Err: %v\n", err)
			return err
		}

		fmt.Fprintf(f, "%s\n\n", localStorage.Value.String())
		fmt.Fprintf(f, "%s\n\n", sessionStorage.Value.String())
		fmt.Fprintf(f, "virtual_dom\n%s\n\n", domData)
		f.Close()
	}
	fmt.Printf("[done] %s\n", url)
	// 	time.Sleep(10000 * time.Second)
	return nil
}

func (b *Browser) HIjack(dirPath string) func(ctx *rod.Hijack) {
	return func(ctx *rod.Hijack) {
		bRequest := ctx.Request
		genName := uuid.New().String()

		indexPath := filepath.Join(dirPath, "index.txt")
		fIndex, err := os.OpenFile(indexPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return
		}
		ctx.LoadResponse(&http.Client{Timeout: time.Duration(time.Second * 5)}, true)
		if ctx.Request.Type() == proto.NetworkResourceTypeMedia || ctx.Request.Type() == proto.NetworkResourceTypeStylesheet || ctx.Request.Type() == proto.NetworkResourceTypeImage || ctx.Request.Type() == proto.NetworkResourceTypeFont {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}

		ctx.ContinueRequest(&proto.FetchContinueRequest{})

		fmt.Fprintf(
			fIndex,
			"%s %s %s %d\n",
			genName,
			bRequest.Req().Method,
			bRequest.Req().URL.String(),
			ctx.Response.Payload().ResponseCode,
		)

		if ctx.Request.Type() == proto.NetworkResourceTypeScript {
			hmd5 := md5.Sum([]byte(bRequest.Req().URL.Path))
			fileName := fmt.Sprintf("%x", hmd5)
			p := filepath.Join(dirPath, "resources")
			p = filepath.Join(p, fileName)

			if !FileExists(p) {
				res, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
				if err != nil {
					fmt.Printf("%v", err)
					return
				}

				fmt.Fprintf(res, "%s\n\n", bRequest.Req().URL)
				fmt.Fprintln(res, string(ctx.Response.Payload().Body))
				res.Close()
			}
		}

		cp := filepath.Join(dirPath, genName)
		fPayload, err := os.OpenFile(cp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return
		}
		kurl := ctx.Request.URL().String()
		if kurl[len(kurl)-1] == '/' {
			kurl = kurl[:len(kurl)-1]
		}
		b.data.Add(kurl, genName)

		fmt.Fprintf(
			fPayload,
			"%s\n\n%s\n\n%s\n\n%s\n\n",
			ctx.Request.URL().String(),
			ctx.Request.Req().Method,
			prettyReqHeaders(ctx.Request.Headers()),
			prettyRespHeaders(ctx.Response.Payload().ResponseHeaders),
		)

		if bRequest.Req().PostForm.Encode() != "" {
			fmt.Fprintf(fPayload,
				"post_params: %s\n",
				bRequest.Req().PostForm.Encode(),
			)
		}

		fIndex.Close()
		fPayload.Close()
	}
}

func prettyReqHeaders(headers proto.NetworkHeaders) string {
	var res string
	for key, val := range headers {
		res += fmt.Sprintf("> %s: %s\n", key, val.String())
	}

	return res
}

func prettyRespHeaders(headers []*proto.FetchHeaderEntry) string {
	var res string
	for _, pair := range headers {
		res += fmt.Sprintf("< %s: %s\n", pair.Name, pair.Value)
	}

	return res
}

func prettyReqCookies(cookies []*http.Cookie) string {
	var res string
	for _, cookie := range cookies {
		res += fmt.Sprintf("> %s: %s\n", cookie.Name, cookie.Value)
	}

	return res
}

func FileExists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}


