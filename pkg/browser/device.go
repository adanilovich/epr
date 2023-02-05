package browser

import "github.com/go-rod/rod/lib/devices"

type Device struct {
}

func NewIPadMini(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 3,
			Horizontal: devices.ScreenSize{
				Height: 768,
				Width:  1024,
			},
			Vertical: devices.ScreenSize{
				Height: 1024,
				Width:  768,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "iPad Mini",
		UserAgent:      `Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1`,
	}
}
func NewIPhone6or7or8Plus(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 3,
			Horizontal: devices.ScreenSize{
				Height: 414,
				Width:  736,
			},
			Vertical: devices.ScreenSize{
				Height: 736,
				Width:  414,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "iPhone 6/7/8 Plus",
		UserAgent:      `Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1`,
	}
}
func NewBlackBerryPlayBook(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 2,
			Horizontal: devices.ScreenSize{
				Height: 600,
				Width:  1024,
			},
			Vertical: devices.ScreenSize{
				Height: 1024,
				Width:  600,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "Blackberry PlayBook",
		UserAgent:      `Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML like Gecko) Version/7.2.1.0 Safari/536.2+`,
	}
}

func NewIPadPro(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 2,
			Horizontal: devices.ScreenSize{
				Height: 1024,
				Width:  1366,
			},
			Vertical: devices.ScreenSize{
				Height: 1366,
				Width:  1024,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "iPad Pro",
		UserAgent:      `Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1`,
	}
}
func NewKindlerFireHDX(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 2,
			Horizontal: devices.ScreenSize{
				Height: 800,
				Width:  1280,
			},
			Vertical: devices.ScreenSize{
				Height: 1280,
				Width:  800,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "Kindle Fire HDX",
		UserAgent:      `Mozilla/5.0 (Linux; U; en-us; KFAPWI Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.13 Safari/535.19 Silk-Accelerated=true`,
	}
}
func NewJioPhone2(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 1,
			Horizontal: devices.ScreenSize{
				Height: 240,
				Width:  320,
			},
			Vertical: devices.ScreenSize{
				Height: 320,
				Width:  240,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "JioPhone 2",
		UserAgent:      `Mozilla/5.0 (Mobile; LYF/F300B/LYF-F300B-001-01-15-130718-i;Android; rv:48.0) Gecko/48.0 Firefox/48.0 KAIOS/2.5`,
	}
}
func NewMicrosoftLumia950(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 2,
			Horizontal: devices.ScreenSize{
				Height: 360,
				Width:  640,
			},
			Vertical: devices.ScreenSize{
				Height: 360,
				Width:  640,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "Microsoft Lumia 550",
		UserAgent:      `Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 550) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/14.14263`,
	}
}
func NewMicrosoftLumia520(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 2,
			Horizontal: devices.ScreenSize{
				Height: 360,
				Width:  640,
			},
			Vertical: devices.ScreenSize{
				Height: 360,
				Width:  640,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "Microsoft Lumia 550",
		UserAgent:      `Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 550) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/14.14263`,
	}
}
func NewNokiaLumia950(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 1.5,
			Horizontal: devices.ScreenSize{
				Height: 360,
				Width:  640,
			},
			Vertical: devices.ScreenSize{
				Height: 640,
				Width:  360,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "Microsoft Lumia 950",
		UserAgent:      `Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/14.14263`,
	}
}

func NewNokia9(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 1,
			Horizontal: devices.ScreenSize{
				Height: 480,
				Width:  854,
			},
			Vertical: devices.ScreenSize{
				Height: 854,
				Width:  480,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "Nokia N9",
		UserAgent:      `Mozilla/5.0 (MeeGo; NokiaN9) AppleWebKit/534.13 (KHTML, like Gecko) NokiaBrowser/8.5.0 Mobile Safari/534.13`,
	}
}
func NewIphone4(acceptLang string) devices.Device {
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 2,
			Horizontal: devices.ScreenSize{
				Height: 320,
				Width:  480,
			},
			Vertical: devices.ScreenSize{
				Height: 480,
				Width:  320,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "iPhone 4",
		UserAgent:      `Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53`,
	}
}
func NewIphoneX(acceptLang string) devices.Device {
	//	// IPhoneX device
	return devices.Device{
		Capabilities: []string{
			"touch",
			"mobile",
		},
		Screen: devices.Screen{
			DevicePixelRatio: 3,
			Horizontal: devices.ScreenSize{
				Height: 375,
				Width:  812,
			},
			Vertical: devices.ScreenSize{
				Height: 812,
				Width:  375,
			},
		},
		AcceptLanguage: acceptLang,
		Title:          "iPhone X",
		UserAgent:      `Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1 Edg/86.0.4240.111`,
	}
}
