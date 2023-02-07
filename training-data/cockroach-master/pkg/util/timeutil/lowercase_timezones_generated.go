// Code generated by pkg/util/timeutil/gen/main.go - DO NOT EDIT.
//
// Copyright 2020 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package timeutil

var lowercaseTimezones = map[string]string{
	`africa/abidjan`:                   `Africa/Abidjan`,
	`africa/accra`:                     `Africa/Accra`,
	`africa/addis_ababa`:               `Africa/Addis_Ababa`,
	`africa/algiers`:                   `Africa/Algiers`,
	`africa/asmara`:                    `Africa/Asmara`,
	`africa/asmera`:                    `Africa/Asmera`,
	`africa/bamako`:                    `Africa/Bamako`,
	`africa/bangui`:                    `Africa/Bangui`,
	`africa/banjul`:                    `Africa/Banjul`,
	`africa/bissau`:                    `Africa/Bissau`,
	`africa/blantyre`:                  `Africa/Blantyre`,
	`africa/brazzaville`:               `Africa/Brazzaville`,
	`africa/bujumbura`:                 `Africa/Bujumbura`,
	`africa/cairo`:                     `Africa/Cairo`,
	`africa/casablanca`:                `Africa/Casablanca`,
	`africa/ceuta`:                     `Africa/Ceuta`,
	`africa/conakry`:                   `Africa/Conakry`,
	`africa/dakar`:                     `Africa/Dakar`,
	`africa/dar_es_salaam`:             `Africa/Dar_es_Salaam`,
	`africa/djibouti`:                  `Africa/Djibouti`,
	`africa/douala`:                    `Africa/Douala`,
	`africa/el_aaiun`:                  `Africa/El_Aaiun`,
	`africa/freetown`:                  `Africa/Freetown`,
	`africa/gaborone`:                  `Africa/Gaborone`,
	`africa/harare`:                    `Africa/Harare`,
	`africa/johannesburg`:              `Africa/Johannesburg`,
	`africa/juba`:                      `Africa/Juba`,
	`africa/kampala`:                   `Africa/Kampala`,
	`africa/khartoum`:                  `Africa/Khartoum`,
	`africa/kigali`:                    `Africa/Kigali`,
	`africa/kinshasa`:                  `Africa/Kinshasa`,
	`africa/lagos`:                     `Africa/Lagos`,
	`africa/libreville`:                `Africa/Libreville`,
	`africa/lome`:                      `Africa/Lome`,
	`africa/luanda`:                    `Africa/Luanda`,
	`africa/lubumbashi`:                `Africa/Lubumbashi`,
	`africa/lusaka`:                    `Africa/Lusaka`,
	`africa/malabo`:                    `Africa/Malabo`,
	`africa/maputo`:                    `Africa/Maputo`,
	`africa/maseru`:                    `Africa/Maseru`,
	`africa/mbabane`:                   `Africa/Mbabane`,
	`africa/mogadishu`:                 `Africa/Mogadishu`,
	`africa/monrovia`:                  `Africa/Monrovia`,
	`africa/nairobi`:                   `Africa/Nairobi`,
	`africa/ndjamena`:                  `Africa/Ndjamena`,
	`africa/niamey`:                    `Africa/Niamey`,
	`africa/nouakchott`:                `Africa/Nouakchott`,
	`africa/ouagadougou`:               `Africa/Ouagadougou`,
	`africa/porto-novo`:                `Africa/Porto-Novo`,
	`africa/sao_tome`:                  `Africa/Sao_Tome`,
	`africa/timbuktu`:                  `Africa/Timbuktu`,
	`africa/tripoli`:                   `Africa/Tripoli`,
	`africa/tunis`:                     `Africa/Tunis`,
	`africa/windhoek`:                  `Africa/Windhoek`,
	`america/adak`:                     `America/Adak`,
	`america/anchorage`:                `America/Anchorage`,
	`america/anguilla`:                 `America/Anguilla`,
	`america/antigua`:                  `America/Antigua`,
	`america/araguaina`:                `America/Araguaina`,
	`america/argentina/buenos_aires`:   `America/Argentina/Buenos_Aires`,
	`america/argentina/catamarca`:      `America/Argentina/Catamarca`,
	`america/argentina/comodrivadavia`: `America/Argentina/ComodRivadavia`,
	`america/argentina/cordoba`:        `America/Argentina/Cordoba`,
	`america/argentina/jujuy`:          `America/Argentina/Jujuy`,
	`america/argentina/la_rioja`:       `America/Argentina/La_Rioja`,
	`america/argentina/mendoza`:        `America/Argentina/Mendoza`,
	`america/argentina/rio_gallegos`:   `America/Argentina/Rio_Gallegos`,
	`america/argentina/salta`:          `America/Argentina/Salta`,
	`america/argentina/san_juan`:       `America/Argentina/San_Juan`,
	`america/argentina/san_luis`:       `America/Argentina/San_Luis`,
	`america/argentina/tucuman`:        `America/Argentina/Tucuman`,
	`america/argentina/ushuaia`:        `America/Argentina/Ushuaia`,
	`america/aruba`:                    `America/Aruba`,
	`america/asuncion`:                 `America/Asuncion`,
	`america/atikokan`:                 `America/Atikokan`,
	`america/atka`:                     `America/Atka`,
	`america/bahia`:                    `America/Bahia`,
	`america/bahia_banderas`:           `America/Bahia_Banderas`,
	`america/barbados`:                 `America/Barbados`,
	`america/belem`:                    `America/Belem`,
	`america/belize`:                   `America/Belize`,
	`america/blanc-sablon`:             `America/Blanc-Sablon`,
	`america/boa_vista`:                `America/Boa_Vista`,
	`america/bogota`:                   `America/Bogota`,
	`america/boise`:                    `America/Boise`,
	`america/buenos_aires`:             `America/Buenos_Aires`,
	`america/cambridge_bay`:            `America/Cambridge_Bay`,
	`america/campo_grande`:             `America/Campo_Grande`,
	`america/cancun`:                   `America/Cancun`,
	`america/caracas`:                  `America/Caracas`,
	`america/catamarca`:                `America/Catamarca`,
	`america/cayenne`:                  `America/Cayenne`,
	`america/cayman`:                   `America/Cayman`,
	`america/chicago`:                  `America/Chicago`,
	`america/chihuahua`:                `America/Chihuahua`,
	`america/coral_harbour`:            `America/Coral_Harbour`,
	`america/cordoba`:                  `America/Cordoba`,
	`america/costa_rica`:               `America/Costa_Rica`,
	`america/creston`:                  `America/Creston`,
	`america/cuiaba`:                   `America/Cuiaba`,
	`america/curacao`:                  `America/Curacao`,
	`america/danmarkshavn`:             `America/Danmarkshavn`,
	`america/dawson`:                   `America/Dawson`,
	`america/dawson_creek`:             `America/Dawson_Creek`,
	`america/denver`:                   `America/Denver`,
	`america/detroit`:                  `America/Detroit`,
	`america/dominica`:                 `America/Dominica`,
	`america/edmonton`:                 `America/Edmonton`,
	`america/eirunepe`:                 `America/Eirunepe`,
	`america/el_salvador`:              `America/El_Salvador`,
	`america/ensenada`:                 `America/Ensenada`,
	`america/fort_nelson`:              `America/Fort_Nelson`,
	`america/fort_wayne`:               `America/Fort_Wayne`,
	`america/fortaleza`:                `America/Fortaleza`,
	`america/glace_bay`:                `America/Glace_Bay`,
	`america/godthab`:                  `America/Godthab`,
	`america/goose_bay`:                `America/Goose_Bay`,
	`america/grand_turk`:               `America/Grand_Turk`,
	`america/grenada`:                  `America/Grenada`,
	`america/guadeloupe`:               `America/Guadeloupe`,
	`america/guatemala`:                `America/Guatemala`,
	`america/guayaquil`:                `America/Guayaquil`,
	`america/guyana`:                   `America/Guyana`,
	`america/halifax`:                  `America/Halifax`,
	`america/havana`:                   `America/Havana`,
	`america/hermosillo`:               `America/Hermosillo`,
	`america/indiana/indianapolis`:     `America/Indiana/Indianapolis`,
	`america/indiana/knox`:             `America/Indiana/Knox`,
	`america/indiana/marengo`:          `America/Indiana/Marengo`,
	`america/indiana/petersburg`:       `America/Indiana/Petersburg`,
	`america/indiana/tell_city`:        `America/Indiana/Tell_City`,
	`america/indiana/vevay`:            `America/Indiana/Vevay`,
	`america/indiana/vincennes`:        `America/Indiana/Vincennes`,
	`america/indiana/winamac`:          `America/Indiana/Winamac`,
	`america/indianapolis`:             `America/Indianapolis`,
	`america/inuvik`:                   `America/Inuvik`,
	`america/iqaluit`:                  `America/Iqaluit`,
	`america/jamaica`:                  `America/Jamaica`,
	`america/jujuy`:                    `America/Jujuy`,
	`america/juneau`:                   `America/Juneau`,
	`america/kentucky/louisville`:      `America/Kentucky/Louisville`,
	`america/kentucky/monticello`:      `America/Kentucky/Monticello`,
	`america/knox_in`:                  `America/Knox_IN`,
	`america/kralendijk`:               `America/Kralendijk`,
	`america/la_paz`:                   `America/La_Paz`,
	`america/lima`:                     `America/Lima`,
	`america/los_angeles`:              `America/Los_Angeles`,
	`america/louisville`:               `America/Louisville`,
	`america/lower_princes`:            `America/Lower_Princes`,
	`america/maceio`:                   `America/Maceio`,
	`america/managua`:                  `America/Managua`,
	`america/manaus`:                   `America/Manaus`,
	`america/marigot`:                  `America/Marigot`,
	`america/martinique`:               `America/Martinique`,
	`america/matamoros`:                `America/Matamoros`,
	`america/mazatlan`:                 `America/Mazatlan`,
	`america/mendoza`:                  `America/Mendoza`,
	`america/menominee`:                `America/Menominee`,
	`america/merida`:                   `America/Merida`,
	`america/metlakatla`:               `America/Metlakatla`,
	`america/mexico_city`:              `America/Mexico_City`,
	`america/miquelon`:                 `America/Miquelon`,
	`america/moncton`:                  `America/Moncton`,
	`america/monterrey`:                `America/Monterrey`,
	`america/montevideo`:               `America/Montevideo`,
	`america/montreal`:                 `America/Montreal`,
	`america/montserrat`:               `America/Montserrat`,
	`america/nassau`:                   `America/Nassau`,
	`america/new_york`:                 `America/New_York`,
	`america/nipigon`:                  `America/Nipigon`,
	`america/nome`:                     `America/Nome`,
	`america/noronha`:                  `America/Noronha`,
	`america/north_dakota/beulah`:      `America/North_Dakota/Beulah`,
	`america/north_dakota/center`:      `America/North_Dakota/Center`,
	`america/north_dakota/new_salem`:   `America/North_Dakota/New_Salem`,
	`america/nuuk`:                     `America/Nuuk`,
	`america/ojinaga`:                  `America/Ojinaga`,
	`america/panama`:                   `America/Panama`,
	`america/pangnirtung`:              `America/Pangnirtung`,
	`america/paramaribo`:               `America/Paramaribo`,
	`america/phoenix`:                  `America/Phoenix`,
	`america/port-au-prince`:           `America/Port-au-Prince`,
	`america/port_of_spain`:            `America/Port_of_Spain`,
	`america/porto_acre`:               `America/Porto_Acre`,
	`america/porto_velho`:              `America/Porto_Velho`,
	`america/puerto_rico`:              `America/Puerto_Rico`,
	`america/punta_arenas`:             `America/Punta_Arenas`,
	`america/rainy_river`:              `America/Rainy_River`,
	`america/rankin_inlet`:             `America/Rankin_Inlet`,
	`america/recife`:                   `America/Recife`,
	`america/regina`:                   `America/Regina`,
	`america/resolute`:                 `America/Resolute`,
	`america/rio_branco`:               `America/Rio_Branco`,
	`america/rosario`:                  `America/Rosario`,
	`america/santa_isabel`:             `America/Santa_Isabel`,
	`america/santarem`:                 `America/Santarem`,
	`america/santiago`:                 `America/Santiago`,
	`america/santo_domingo`:            `America/Santo_Domingo`,
	`america/sao_paulo`:                `America/Sao_Paulo`,
	`america/scoresbysund`:             `America/Scoresbysund`,
	`america/shiprock`:                 `America/Shiprock`,
	`america/sitka`:                    `America/Sitka`,
	`america/st_barthelemy`:            `America/St_Barthelemy`,
	`america/st_johns`:                 `America/St_Johns`,
	`america/st_kitts`:                 `America/St_Kitts`,
	`america/st_lucia`:                 `America/St_Lucia`,
	`america/st_thomas`:                `America/St_Thomas`,
	`america/st_vincent`:               `America/St_Vincent`,
	`america/swift_current`:            `America/Swift_Current`,
	`america/tegucigalpa`:              `America/Tegucigalpa`,
	`america/thule`:                    `America/Thule`,
	`america/thunder_bay`:              `America/Thunder_Bay`,
	`america/tijuana`:                  `America/Tijuana`,
	`america/toronto`:                  `America/Toronto`,
	`america/tortola`:                  `America/Tortola`,
	`america/vancouver`:                `America/Vancouver`,
	`america/virgin`:                   `America/Virgin`,
	`america/whitehorse`:               `America/Whitehorse`,
	`america/winnipeg`:                 `America/Winnipeg`,
	`america/yakutat`:                  `America/Yakutat`,
	`america/yellowknife`:              `America/Yellowknife`,
	`antarctica/casey`:                 `Antarctica/Casey`,
	`antarctica/davis`:                 `Antarctica/Davis`,
	`antarctica/dumontdurville`:        `Antarctica/DumontDUrville`,
	`antarctica/macquarie`:             `Antarctica/Macquarie`,
	`antarctica/mawson`:                `Antarctica/Mawson`,
	`antarctica/mcmurdo`:               `Antarctica/McMurdo`,
	`antarctica/palmer`:                `Antarctica/Palmer`,
	`antarctica/rothera`:               `Antarctica/Rothera`,
	`antarctica/south_pole`:            `Antarctica/South_Pole`,
	`antarctica/syowa`:                 `Antarctica/Syowa`,
	`antarctica/troll`:                 `Antarctica/Troll`,
	`antarctica/vostok`:                `Antarctica/Vostok`,
	`arctic/longyearbyen`:              `Arctic/Longyearbyen`,
	`asia/aden`:                        `Asia/Aden`,
	`asia/almaty`:                      `Asia/Almaty`,
	`asia/amman`:                       `Asia/Amman`,
	`asia/anadyr`:                      `Asia/Anadyr`,
	`asia/aqtau`:                       `Asia/Aqtau`,
	`asia/aqtobe`:                      `Asia/Aqtobe`,
	`asia/ashgabat`:                    `Asia/Ashgabat`,
	`asia/ashkhabad`:                   `Asia/Ashkhabad`,
	`asia/atyrau`:                      `Asia/Atyrau`,
	`asia/baghdad`:                     `Asia/Baghdad`,
	`asia/bahrain`:                     `Asia/Bahrain`,
	`asia/baku`:                        `Asia/Baku`,
	`asia/bangkok`:                     `Asia/Bangkok`,
	`asia/barnaul`:                     `Asia/Barnaul`,
	`asia/beirut`:                      `Asia/Beirut`,
	`asia/bishkek`:                     `Asia/Bishkek`,
	`asia/brunei`:                      `Asia/Brunei`,
	`asia/calcutta`:                    `Asia/Calcutta`,
	`asia/chita`:                       `Asia/Chita`,
	`asia/choibalsan`:                  `Asia/Choibalsan`,
	`asia/chongqing`:                   `Asia/Chongqing`,
	`asia/chungking`:                   `Asia/Chungking`,
	`asia/colombo`:                     `Asia/Colombo`,
	`asia/dacca`:                       `Asia/Dacca`,
	`asia/damascus`:                    `Asia/Damascus`,
	`asia/dhaka`:                       `Asia/Dhaka`,
	`asia/dili`:                        `Asia/Dili`,
	`asia/dubai`:                       `Asia/Dubai`,
	`asia/dushanbe`:                    `Asia/Dushanbe`,
	`asia/famagusta`:                   `Asia/Famagusta`,
	`asia/gaza`:                        `Asia/Gaza`,
	`asia/harbin`:                      `Asia/Harbin`,
	`asia/hebron`:                      `Asia/Hebron`,
	`asia/ho_chi_minh`:                 `Asia/Ho_Chi_Minh`,
	`asia/hong_kong`:                   `Asia/Hong_Kong`,
	`asia/hovd`:                        `Asia/Hovd`,
	`asia/irkutsk`:                     `Asia/Irkutsk`,
	`asia/istanbul`:                    `Asia/Istanbul`,
	`asia/jakarta`:                     `Asia/Jakarta`,
	`asia/jayapura`:                    `Asia/Jayapura`,
	`asia/jerusalem`:                   `Asia/Jerusalem`,
	`asia/kabul`:                       `Asia/Kabul`,
	`asia/kamchatka`:                   `Asia/Kamchatka`,
	`asia/karachi`:                     `Asia/Karachi`,
	`asia/kashgar`:                     `Asia/Kashgar`,
	`asia/kathmandu`:                   `Asia/Kathmandu`,
	`asia/katmandu`:                    `Asia/Katmandu`,
	`asia/khandyga`:                    `Asia/Khandyga`,
	`asia/kolkata`:                     `Asia/Kolkata`,
	`asia/krasnoyarsk`:                 `Asia/Krasnoyarsk`,
	`asia/kuala_lumpur`:                `Asia/Kuala_Lumpur`,
	`asia/kuching`:                     `Asia/Kuching`,
	`asia/kuwait`:                      `Asia/Kuwait`,
	`asia/macao`:                       `Asia/Macao`,
	`asia/macau`:                       `Asia/Macau`,
	`asia/magadan`:                     `Asia/Magadan`,
	`asia/makassar`:                    `Asia/Makassar`,
	`asia/manila`:                      `Asia/Manila`,
	`asia/muscat`:                      `Asia/Muscat`,
	`asia/nicosia`:                     `Asia/Nicosia`,
	`asia/novokuznetsk`:                `Asia/Novokuznetsk`,
	`asia/novosibirsk`:                 `Asia/Novosibirsk`,
	`asia/omsk`:                        `Asia/Omsk`,
	`asia/oral`:                        `Asia/Oral`,
	`asia/phnom_penh`:                  `Asia/Phnom_Penh`,
	`asia/pontianak`:                   `Asia/Pontianak`,
	`asia/pyongyang`:                   `Asia/Pyongyang`,
	`asia/qatar`:                       `Asia/Qatar`,
	`asia/qostanay`:                    `Asia/Qostanay`,
	`asia/qyzylorda`:                   `Asia/Qyzylorda`,
	`asia/rangoon`:                     `Asia/Rangoon`,
	`asia/riyadh`:                      `Asia/Riyadh`,
	`asia/saigon`:                      `Asia/Saigon`,
	`asia/sakhalin`:                    `Asia/Sakhalin`,
	`asia/samarkand`:                   `Asia/Samarkand`,
	`asia/seoul`:                       `Asia/Seoul`,
	`asia/shanghai`:                    `Asia/Shanghai`,
	`asia/singapore`:                   `Asia/Singapore`,
	`asia/srednekolymsk`:               `Asia/Srednekolymsk`,
	`asia/taipei`:                      `Asia/Taipei`,
	`asia/tashkent`:                    `Asia/Tashkent`,
	`asia/tbilisi`:                     `Asia/Tbilisi`,
	`asia/tehran`:                      `Asia/Tehran`,
	`asia/tel_aviv`:                    `Asia/Tel_Aviv`,
	`asia/thimbu`:                      `Asia/Thimbu`,
	`asia/thimphu`:                     `Asia/Thimphu`,
	`asia/tokyo`:                       `Asia/Tokyo`,
	`asia/tomsk`:                       `Asia/Tomsk`,
	`asia/ujung_pandang`:               `Asia/Ujung_Pandang`,
	`asia/ulaanbaatar`:                 `Asia/Ulaanbaatar`,
	`asia/ulan_bator`:                  `Asia/Ulan_Bator`,
	`asia/urumqi`:                      `Asia/Urumqi`,
	`asia/ust-nera`:                    `Asia/Ust-Nera`,
	`asia/vientiane`:                   `Asia/Vientiane`,
	`asia/vladivostok`:                 `Asia/Vladivostok`,
	`asia/yakutsk`:                     `Asia/Yakutsk`,
	`asia/yangon`:                      `Asia/Yangon`,
	`asia/yekaterinburg`:               `Asia/Yekaterinburg`,
	`asia/yerevan`:                     `Asia/Yerevan`,
	`atlantic/azores`:                  `Atlantic/Azores`,
	`atlantic/bermuda`:                 `Atlantic/Bermuda`,
	`atlantic/canary`:                  `Atlantic/Canary`,
	`atlantic/cape_verde`:              `Atlantic/Cape_Verde`,
	`atlantic/faeroe`:                  `Atlantic/Faeroe`,
	`atlantic/faroe`:                   `Atlantic/Faroe`,
	`atlantic/jan_mayen`:               `Atlantic/Jan_Mayen`,
	`atlantic/madeira`:                 `Atlantic/Madeira`,
	`atlantic/reykjavik`:               `Atlantic/Reykjavik`,
	`atlantic/south_georgia`:           `Atlantic/South_Georgia`,
	`atlantic/st_helena`:               `Atlantic/St_Helena`,
	`atlantic/stanley`:                 `Atlantic/Stanley`,
	`australia/act`:                    `Australia/ACT`,
	`australia/adelaide`:               `Australia/Adelaide`,
	`australia/brisbane`:               `Australia/Brisbane`,
	`australia/broken_hill`:            `Australia/Broken_Hill`,
	`australia/canberra`:               `Australia/Canberra`,
	`australia/currie`:                 `Australia/Currie`,
	`australia/darwin`:                 `Australia/Darwin`,
	`australia/eucla`:                  `Australia/Eucla`,
	`australia/hobart`:                 `Australia/Hobart`,
	`australia/lhi`:                    `Australia/LHI`,
	`australia/lindeman`:               `Australia/Lindeman`,
	`australia/lord_howe`:              `Australia/Lord_Howe`,
	`australia/melbourne`:              `Australia/Melbourne`,
	`australia/nsw`:                    `Australia/NSW`,
	`australia/north`:                  `Australia/North`,
	`australia/perth`:                  `Australia/Perth`,
	`australia/queensland`:             `Australia/Queensland`,
	`australia/south`:                  `Australia/South`,
	`australia/sydney`:                 `Australia/Sydney`,
	`australia/tasmania`:               `Australia/Tasmania`,
	`australia/victoria`:               `Australia/Victoria`,
	`australia/west`:                   `Australia/West`,
	`australia/yancowinna`:             `Australia/Yancowinna`,
	`brazil/acre`:                      `Brazil/Acre`,
	`brazil/denoronha`:                 `Brazil/DeNoronha`,
	`brazil/east`:                      `Brazil/East`,
	`brazil/west`:                      `Brazil/West`,
	`cet`:                              `CET`,
	`cst6cdt`:                          `CST6CDT`,
	`canada/atlantic`:                  `Canada/Atlantic`,
	`canada/central`:                   `Canada/Central`,
	`canada/eastern`:                   `Canada/Eastern`,
	`canada/mountain`:                  `Canada/Mountain`,
	`canada/newfoundland`:              `Canada/Newfoundland`,
	`canada/pacific`:                   `Canada/Pacific`,
	`canada/saskatchewan`:              `Canada/Saskatchewan`,
	`canada/yukon`:                     `Canada/Yukon`,
	`chile/continental`:                `Chile/Continental`,
	`chile/easterisland`:               `Chile/EasterIsland`,
	`cuba`:                             `Cuba`,
	`eet`:                              `EET`,
	`est`:                              `EST`,
	`est5edt`:                          `EST5EDT`,
	`egypt`:                            `Egypt`,
	`eire`:                             `Eire`,
	`etc/gmt`:                          `Etc/GMT`,
	`etc/gmt+0`:                        `Etc/GMT+0`,
	`etc/gmt+1`:                        `Etc/GMT+1`,
	`etc/gmt+10`:                       `Etc/GMT+10`,
	`etc/gmt+11`:                       `Etc/GMT+11`,
	`etc/gmt+12`:                       `Etc/GMT+12`,
	`etc/gmt+2`:                        `Etc/GMT+2`,
	`etc/gmt+3`:                        `Etc/GMT+3`,
	`etc/gmt+4`:                        `Etc/GMT+4`,
	`etc/gmt+5`:                        `Etc/GMT+5`,
	`etc/gmt+6`:                        `Etc/GMT+6`,
	`etc/gmt+7`:                        `Etc/GMT+7`,
	`etc/gmt+8`:                        `Etc/GMT+8`,
	`etc/gmt+9`:                        `Etc/GMT+9`,
	`etc/gmt-0`:                        `Etc/GMT-0`,
	`etc/gmt-1`:                        `Etc/GMT-1`,
	`etc/gmt-10`:                       `Etc/GMT-10`,
	`etc/gmt-11`:                       `Etc/GMT-11`,
	`etc/gmt-12`:                       `Etc/GMT-12`,
	`etc/gmt-13`:                       `Etc/GMT-13`,
	`etc/gmt-14`:                       `Etc/GMT-14`,
	`etc/gmt-2`:                        `Etc/GMT-2`,
	`etc/gmt-3`:                        `Etc/GMT-3`,
	`etc/gmt-4`:                        `Etc/GMT-4`,
	`etc/gmt-5`:                        `Etc/GMT-5`,
	`etc/gmt-6`:                        `Etc/GMT-6`,
	`etc/gmt-7`:                        `Etc/GMT-7`,
	`etc/gmt-8`:                        `Etc/GMT-8`,
	`etc/gmt-9`:                        `Etc/GMT-9`,
	`etc/gmt0`:                         `Etc/GMT0`,
	`etc/greenwich`:                    `Etc/Greenwich`,
	`etc/uct`:                          `Etc/UCT`,
	`etc/utc`:                          `Etc/UTC`,
	`etc/universal`:                    `Etc/Universal`,
	`etc/zulu`:                         `Etc/Zulu`,
	`europe/amsterdam`:                 `Europe/Amsterdam`,
	`europe/andorra`:                   `Europe/Andorra`,
	`europe/astrakhan`:                 `Europe/Astrakhan`,
	`europe/athens`:                    `Europe/Athens`,
	`europe/belfast`:                   `Europe/Belfast`,
	`europe/belgrade`:                  `Europe/Belgrade`,
	`europe/berlin`:                    `Europe/Berlin`,
	`europe/bratislava`:                `Europe/Bratislava`,
	`europe/brussels`:                  `Europe/Brussels`,
	`europe/bucharest`:                 `Europe/Bucharest`,
	`europe/budapest`:                  `Europe/Budapest`,
	`europe/busingen`:                  `Europe/Busingen`,
	`europe/chisinau`:                  `Europe/Chisinau`,
	`europe/copenhagen`:                `Europe/Copenhagen`,
	`europe/dublin`:                    `Europe/Dublin`,
	`europe/gibraltar`:                 `Europe/Gibraltar`,
	`europe/guernsey`:                  `Europe/Guernsey`,
	`europe/helsinki`:                  `Europe/Helsinki`,
	`europe/isle_of_man`:               `Europe/Isle_of_Man`,
	`europe/istanbul`:                  `Europe/Istanbul`,
	`europe/jersey`:                    `Europe/Jersey`,
	`europe/kaliningrad`:               `Europe/Kaliningrad`,
	`europe/kiev`:                      `Europe/Kiev`,
	`europe/kirov`:                     `Europe/Kirov`,
	`europe/lisbon`:                    `Europe/Lisbon`,
	`europe/ljubljana`:                 `Europe/Ljubljana`,
	`europe/london`:                    `Europe/London`,
	`europe/luxembourg`:                `Europe/Luxembourg`,
	`europe/madrid`:                    `Europe/Madrid`,
	`europe/malta`:                     `Europe/Malta`,
	`europe/mariehamn`:                 `Europe/Mariehamn`,
	`europe/minsk`:                     `Europe/Minsk`,
	`europe/monaco`:                    `Europe/Monaco`,
	`europe/moscow`:                    `Europe/Moscow`,
	`europe/nicosia`:                   `Europe/Nicosia`,
	`europe/oslo`:                      `Europe/Oslo`,
	`europe/paris`:                     `Europe/Paris`,
	`europe/podgorica`:                 `Europe/Podgorica`,
	`europe/prague`:                    `Europe/Prague`,
	`europe/riga`:                      `Europe/Riga`,
	`europe/rome`:                      `Europe/Rome`,
	`europe/samara`:                    `Europe/Samara`,
	`europe/san_marino`:                `Europe/San_Marino`,
	`europe/sarajevo`:                  `Europe/Sarajevo`,
	`europe/saratov`:                   `Europe/Saratov`,
	`europe/simferopol`:                `Europe/Simferopol`,
	`europe/skopje`:                    `Europe/Skopje`,
	`europe/sofia`:                     `Europe/Sofia`,
	`europe/stockholm`:                 `Europe/Stockholm`,
	`europe/tallinn`:                   `Europe/Tallinn`,
	`europe/tirane`:                    `Europe/Tirane`,
	`europe/tiraspol`:                  `Europe/Tiraspol`,
	`europe/ulyanovsk`:                 `Europe/Ulyanovsk`,
	`europe/uzhgorod`:                  `Europe/Uzhgorod`,
	`europe/vaduz`:                     `Europe/Vaduz`,
	`europe/vatican`:                   `Europe/Vatican`,
	`europe/vienna`:                    `Europe/Vienna`,
	`europe/vilnius`:                   `Europe/Vilnius`,
	`europe/volgograd`:                 `Europe/Volgograd`,
	`europe/warsaw`:                    `Europe/Warsaw`,
	`europe/zagreb`:                    `Europe/Zagreb`,
	`europe/zaporozhye`:                `Europe/Zaporozhye`,
	`europe/zurich`:                    `Europe/Zurich`,
	`factory`:                          `Factory`,
	`gb`:                               `GB`,
	`gb-eire`:                          `GB-Eire`,
	`gmt`:                              `GMT`,
	`gmt+0`:                            `GMT+0`,
	`gmt-0`:                            `GMT-0`,
	`gmt0`:                             `GMT0`,
	`greenwich`:                        `Greenwich`,
	`hst`:                              `HST`,
	`hongkong`:                         `Hongkong`,
	`iceland`:                          `Iceland`,
	`indian/antananarivo`:              `Indian/Antananarivo`,
	`indian/chagos`:                    `Indian/Chagos`,
	`indian/christmas`:                 `Indian/Christmas`,
	`indian/cocos`:                     `Indian/Cocos`,
	`indian/comoro`:                    `Indian/Comoro`,
	`indian/kerguelen`:                 `Indian/Kerguelen`,
	`indian/mahe`:                      `Indian/Mahe`,
	`indian/maldives`:                  `Indian/Maldives`,
	`indian/mauritius`:                 `Indian/Mauritius`,
	`indian/mayotte`:                   `Indian/Mayotte`,
	`indian/reunion`:                   `Indian/Reunion`,
	`iran`:                             `Iran`,
	`israel`:                           `Israel`,
	`jamaica`:                          `Jamaica`,
	`japan`:                            `Japan`,
	`kwajalein`:                        `Kwajalein`,
	`libya`:                            `Libya`,
	`met`:                              `MET`,
	`mst`:                              `MST`,
	`mst7mdt`:                          `MST7MDT`,
	`mexico/bajanorte`:                 `Mexico/BajaNorte`,
	`mexico/bajasur`:                   `Mexico/BajaSur`,
	`mexico/general`:                   `Mexico/General`,
	`nz`:                               `NZ`,
	`nz-chat`:                          `NZ-CHAT`,
	`navajo`:                           `Navajo`,
	`prc`:                              `PRC`,
	`pst8pdt`:                          `PST8PDT`,
	`pacific/apia`:                     `Pacific/Apia`,
	`pacific/auckland`:                 `Pacific/Auckland`,
	`pacific/bougainville`:             `Pacific/Bougainville`,
	`pacific/chatham`:                  `Pacific/Chatham`,
	`pacific/chuuk`:                    `Pacific/Chuuk`,
	`pacific/easter`:                   `Pacific/Easter`,
	`pacific/efate`:                    `Pacific/Efate`,
	`pacific/enderbury`:                `Pacific/Enderbury`,
	`pacific/fakaofo`:                  `Pacific/Fakaofo`,
	`pacific/fiji`:                     `Pacific/Fiji`,
	`pacific/funafuti`:                 `Pacific/Funafuti`,
	`pacific/galapagos`:                `Pacific/Galapagos`,
	`pacific/gambier`:                  `Pacific/Gambier`,
	`pacific/guadalcanal`:              `Pacific/Guadalcanal`,
	`pacific/guam`:                     `Pacific/Guam`,
	`pacific/honolulu`:                 `Pacific/Honolulu`,
	`pacific/johnston`:                 `Pacific/Johnston`,
	`pacific/kanton`:                   `Pacific/Kanton`,
	`pacific/kiritimati`:               `Pacific/Kiritimati`,
	`pacific/kosrae`:                   `Pacific/Kosrae`,
	`pacific/kwajalein`:                `Pacific/Kwajalein`,
	`pacific/majuro`:                   `Pacific/Majuro`,
	`pacific/marquesas`:                `Pacific/Marquesas`,
	`pacific/midway`:                   `Pacific/Midway`,
	`pacific/nauru`:                    `Pacific/Nauru`,
	`pacific/niue`:                     `Pacific/Niue`,
	`pacific/norfolk`:                  `Pacific/Norfolk`,
	`pacific/noumea`:                   `Pacific/Noumea`,
	`pacific/pago_pago`:                `Pacific/Pago_Pago`,
	`pacific/palau`:                    `Pacific/Palau`,
	`pacific/pitcairn`:                 `Pacific/Pitcairn`,
	`pacific/pohnpei`:                  `Pacific/Pohnpei`,
	`pacific/ponape`:                   `Pacific/Ponape`,
	`pacific/port_moresby`:             `Pacific/Port_Moresby`,
	`pacific/rarotonga`:                `Pacific/Rarotonga`,
	`pacific/saipan`:                   `Pacific/Saipan`,
	`pacific/samoa`:                    `Pacific/Samoa`,
	`pacific/tahiti`:                   `Pacific/Tahiti`,
	`pacific/tarawa`:                   `Pacific/Tarawa`,
	`pacific/tongatapu`:                `Pacific/Tongatapu`,
	`pacific/truk`:                     `Pacific/Truk`,
	`pacific/wake`:                     `Pacific/Wake`,
	`pacific/wallis`:                   `Pacific/Wallis`,
	`pacific/yap`:                      `Pacific/Yap`,
	`poland`:                           `Poland`,
	`portugal`:                         `Portugal`,
	`roc`:                              `ROC`,
	`rok`:                              `ROK`,
	`singapore`:                        `Singapore`,
	`turkey`:                           `Turkey`,
	`uct`:                              `UCT`,
	`us/alaska`:                        `US/Alaska`,
	`us/aleutian`:                      `US/Aleutian`,
	`us/arizona`:                       `US/Arizona`,
	`us/central`:                       `US/Central`,
	`us/east-indiana`:                  `US/East-Indiana`,
	`us/eastern`:                       `US/Eastern`,
	`us/hawaii`:                        `US/Hawaii`,
	`us/indiana-starke`:                `US/Indiana-Starke`,
	`us/michigan`:                      `US/Michigan`,
	`us/mountain`:                      `US/Mountain`,
	`us/pacific`:                       `US/Pacific`,
	`us/samoa`:                         `US/Samoa`,
	`utc`:                              `UTC`,
	`universal`:                        `Universal`,
	`w-su`:                             `W-SU`,
	`wet`:                              `WET`,
	`zulu`:                             `Zulu`,
}
