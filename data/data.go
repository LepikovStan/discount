package data

var d = map[string]map[string][]string{
	"MoonyXLBoys": {
		"akusherstvo": []string{
			"https://www.akusherstvo.ru/magaz.php?action=show_tovar&tovar_id=20430",
		},
		"ozon": []string{
			"https://www.ozon.ru/context/detail/id/135608220/",
		},
		"ulmart": []string{
			"https://m.ulmart.ru/goods/3801519",
		},
		"dochkisinochki": []string{
			"http://www.dochkisinochki.ru/icatalog/products/2693780/",
		},
		"mytoys": []string{
			"http://www.mytoys.ru/product/4659615",
		},
		"helptomama": []string{
			"https://helptomama.ru/catalog/yaponskie-podguzniki/yaponskie_podguzniki_trusiki_moony_dlya_malchikov_xl_12_17kg_38sht/",
		},
	},
}

func Get(name string) map[string][]string {
	return d[name]
}
