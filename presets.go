package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
	"austria":      PresetLocations{"austria", "%C3%B6sterreich", "vienna", "wien", "linz", "salzburg", "graz", "innsbruck", "klagenfurt", "wels", "dornbirn"},
	"finland":      PresetLocations{"finland", "suomi", "helsinki", "tampere", "oulu", "espoo", "vantaa", "turku"},
	"sweden":       PresetLocations{"sweden", "sverige", "stockholm", "malm%C3%B6", "uppsala", "g%C3%B6teborg", "gothenburg", "lund"},
	"norway":       PresetLocations{"norway", "norge", "oslo", "bergen", "trondheim", "sandnes"},
	"germany":      PresetLocations{"germany", "deutschland", "berlin", "frankfurt", "munich", "m%C3%BCnchen", "hamburg", "cologne", "k%C3%B6ln"},
	"netherlands":  PresetLocations{"netherlands", "nederland", "amsterdam", "rotterdam", "hague", "utrecht", "holland", "delft"},
	"ukraine":      PresetLocations{"ukraine", "kiev", "kyiv", "kharkiv", "dnipro", "odesa", "donetsk", "zaporizhia"},
	"japan":        PresetLocations{"japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
	"russia":       PresetLocations{"russia", "moscow", "saint%2Bpetersburg", "novosibirsk", "yekaterinburg", "nizhny%2Bnovgorod", "samara", "omsk", "kazan", "chelyabinsk", "rostov-on-don", "ufa", "volgograd"},
	"estonia":      PresetLocations{"estonia", "eesti", "tallinn", "tartu", "narva", "p%C3%A4rnu", "haapsalu"},
	"denmark":      PresetLocations{"denmark", "danmark", "copenhagen", "aarhus", "odense", "aalborg"},
	"portugal":     PresetLocations{"portugal", "lisbon", "lisboa", "braga", "porto", "aveiro", "coimbra", "funchal", "madeira"},
	"france":       PresetLocations{"france", "paris", "marseille", "lyon", "toulouse", "nice", "nantes", "strasbourg", "montpellier", "bordeaux", "lille", "rennes", "reims"},
	"spain":        PresetLocations{"spain", "espa%C3%B1a", "madrid", "barcelona", "valencia", "seville", "sevilla", "zaragoza", "malaga", "murcia", "palma", "bilbao", "alicante", "cordoba"},
	"italy":        PresetLocations{"italy", "italia", "rome", "roma", "milan", "naples", "napoli", "turin", "torino", "palermo", "genoa", "genova", "bologna", "florence", "firenze", "bari", "catania", "venice", "verona"},
	"uk":           PresetLocations{"uk", "london", "birmingham", "leeds", "glasgow", "sheffield", "bradford", "manchester", "edinburgh", "liverpool", "bristol", "cardiff", "belfast", "leicester", "wakefield", "coventry", "nottingham", "newcastle"},
	"croatia":      PresetLocations{"croatia", "hrvatska", "zagreb", "split", "rijeka", "osijek", "zadar", "pula"},
	"worldwide":    PresetLocations{},
	"china":        PresetLocations{"china", "%E4%B8%AD%E5%9B%BD", "guangzhou", "shanghai", "beijing", "hangzhou", "hong%2Bkong"},
	"india":        PresetLocations{"india", "mumbai", "delhi", "bangalore", "hyderabad", "ahmedabad", "chennai", "kolkata", "jaipur"},
	"israel":       PresetLocations{"israel", "tel%2Baviv", "jerusalem", "beer%2Bsheva", "beersheva", "netanya", "ramat%2Bgan", "haifa", "herzliya", "rishon"},
	"indonesia":    PresetLocations{"indonesia", "jakarta", "surabaya", "bandung", "medan", "bekasi", "semarang", "tangerang", "depok", "makassar", "palembang"},
	"pakistan":     PresetLocations{"pakistan", "karachi", "lahore", "faisalabad", "rawalpindi", "peshawar", "islamabad"},
	"brazil":       PresetLocations{"brazil", "brasil", "s%C3%A3o%20paulo", "bras%C3%ADlia", "salvador", "fortaleza", "belo%20horizonte", "manaus", "curitiba", "recife", "porto%20alegre"},
	"nigeria":      PresetLocations{"nigeria", "lagos", "kano", "ibadan", "benin%20city", "port%20harcourt", "jos", "ilorin"},
	"bangladesh":   PresetLocations{"bangladesh", "dhaka", "chittagong", "khulna", "rajshahi", "barisal", "sylhet", "rangpur", "comilla", "gazipur"},
	"mexico":       PresetLocations{"mexico", "mexico%20city", "guadalajara", "puebla", "tijuana", "mexicali", "monterrey", "hermosillo", "zapopan", "ciudad%20juarez", "chihuahua", "aguascalientes", "mx"},
	"philippines":  PresetLocations{"philippines", "pilipinas", "quezon", "manila", "davao", "caloocan", "cebu", "zamboanga", "bohol", "pasig", "bacolod", "makati", "baguio"},
	"luxembourg":   PresetLocations{"luxembourg", "esch-sur-alzette", "differdange", "dudelange", "ettelbruck", "diekirch", "wiltz", "echternach", "rumelange", "grevenmacher", "bertrange", "mamer", "capellen", "strassen", "diekirch"},
	"egypt":        PresetLocations{"egypt", "cairo", "alexandria", "giza", "port%2Bsaid", "suez", "luxor", "el%2Bmahalla", "asyut", "al%2Bmansurah", "tanda"},
	"ethiopia":     PresetLocations{"ethiopia", "addis%2Bababa", "gondar", "adama", "hawassa", "bahir%2Bdar"},
	"vietnam":      PresetLocations{"vietnam", "viet%2Bnam", "ho%2Bchi%2Bminh", "hanoi", "ha%2Bnoi", "hai%2Bphong", "da%2Bnang", "can%2Btho", "bien%2Bhoa", "nha%Btrang", "vinh"},
	"iran":         PresetLocations{"iran", "tehran", "mashhad", "isfahan", "esfahan", "karaj", "shiraz", "tabriz", "qom", "ahvaz", "ahwaz", "kermanshah", "urmia", "rasht", "kerman"},
	"congo":        PresetLocations{"congo", "drc", "kinshasa", "lubumbashi", "bukavu", "kananga", "likasi", "mbujimayi"},
	"turkey":       PresetLocations{"turkey", "turkiye", "istanbul", "ankara", "izmir", "bursa", "adana", "gaziantep", "konya", "antalya", "kayseri", "mersin", "eskisehir", "samsun", "denizli", "malatya"},
	"thailand":     PresetLocations{"thailand", "bangkok", "nonthaburi", "nakhon", "phuket", "pattaya", "chiang%2Bmai"},
	"south africa": PresetLocations{"south%2Bafrica", "south%2Bafrica", "johannesburg", "cape%2Btown", "rsa", "durban", "port%2Belizabeth", "pretoria", "nelspruit"},
	"myanmar":      PresetLocations{"myanmar", "burma", "yangon", "rangoon", "mandalay", "nay%2Bpyi%2Btaw", "taunggyi", "bago", "mawlamyine"},
	"tanzania":     PresetLocations{"tanzania", "dar%2Bes%2Bsalaam", "mwanza", "arusha", "dodoma", "mbeya", "morogoro", "tanga", "kilimanjaro"},
	"south korea":  PresetLocations{"south%2Bkorea", "ROK", "korea", "seoul", "busan", "incheon", "daegu", "daejeon", "gwangju", "%EB%8C%80%ED%95%9C%EB%AF%BC%EA%B5%AD", "%EC%84%9C%EC%9A%B8", "%EC%84%9C%EC%9A%B8%EC%8B%9C"},
	"colombia":     PresetLocations{"colombia", "bogota", "medellin", "cali", "barranquilla", "cartagena", "cucuta", "bucaramanga", "ibague", "soledad", "pereira", "santa%2Bmarta"},
	"kenya":        PresetLocations{"kenya", "nairobi", "mombasa", "kisumu", "nakuru", "eldoret", "kisii"},
	"argentina":    PresetLocations{"argentina", "buenos%2Baires", "cordoba", "rosario", "mendoza", "la%2Bplata", "tucuman", "mar%2Bdel%2Bplata", "salta", "resistencia"},
	"algeria":      PresetLocations{"algeria", "algiers", "oran", "constantine", "annaba", "blida", "batna", "djelfa", "setif", "sidi%2Bbel%2Babbes", "biskra", "tiaret", "relizane", "mostaganem", "tlemcen", "chlef", "jijel"},
	"sudan":        PresetLocations{"sudan", "khartoum", "omdurman", "kassala"},
	"poland":       PresetLocations{"poland", "polska", "warsaw", "krakow", "lodz", "wroclaw", "poznan", "gdansk", "szczecin", "bydgoszcz", "lublin", "katowice", "bialystok", "opole"},
	"canada":       PresetLocations{"canada", "ottawa", "edmonton", "winnipeg", "vancouver", "toronto", "quebec", "montreal", "mississauga", "calgary"},
	"australia":    PresetLocations{"australia", "sydney", "melbourne", "brisbane", "perth", "adelaide", "canberra", "hobart", "darwin"},
	"new zealand":  PresetLocations{"new%2Bzealand", "auckland", "wellington", "christchurch", "hamilton", "tauranga", "napier-hastings", "dunedin", "palmerston%2Bnorth", "nelson", "rotorua", "whangarei", "new%2Bplymouth", "invercargill", "whanganui", "gisborne"},
	"belgium":      PresetLocations{"belgium", "antwerp", "ghent", "charleroi", "liege", "brussels", "belgique"},
	"greece":       PresetLocations{"greece", "%CE%95%CE%BB%CE%BB%CE%AC%CE%B4%CE%B1", "athens", "thessaloniki", "patras", "heraklion", "larissa", "volos", "rhodes", "ioannina", "chania", "crete"},
	"peru":         PresetLocations{"peru", "lima", "cusco", "cuzco", "ica", "arequipa", "trujillo", "chiclayo", "huancayo", "piura", "chimbote", "iquitos", "juliaca", "cajamarca"},
	"hungary":      PresetLocations{"hungary", "magyarorsz%C3%A1g", "budapest", "szeged", "miskolc"},
	"albania":      PresetLocations{"albania", "tirana", "durres", "vlore", "elbasan", "shkoder"},
	"uganda":       PresetLocations{"uganda", "kampala", "mbarara", "mukono", "jinja", "arua", "gulu", "masaka"},
	"zambia":       PresetLocations{"zambia", "lusaka", "kitwe", "ndola"},
	"sri lanka":    PresetLocations{"sri%2Blanka","nugeegoda", "balangoda", "mawathagama", "ratnapura","medawachchiya", "poonewa", "colombo", "moratuwa", "negombo", "galle", "jaffna", "kelaniya", "ambalangoda", "panadura", "wadduwa", "moronthuduwa", "minuwangoda","kegalle","nuwraeliya","kirulapone","pitigala","alawwa","kurunegala","polgahawela","kuruwita","nugegoda","kalutara","horana","anuradhapura", "nivithigala", "kandy", "delgoda" ,"kalubowila", "homagama", "puthlam", "trincomalee", "kaduwela", "pitipana", "padukka", "akarawita", "bambalapitiya","batugampola", "hokandara", "kiriwattuduwa", "kolonnawa", "madapatha", "meegoda", "mullegama","matara","weligama", "pamunugama","avissawella","tangalle","beliatta","tissamaharama","ambalantota","hambantota","weeraketiya","wattala","jaela","meegamuwa"},
	"singapore":    PresetLocations{"singapore", "tampines","pasir%2Bris","bishan"},
	"latvia":       PresetLocations{"latvia", "latvija", "riga", "r%C4%ABga", "kuldiga", "kuld%C4%ABga", "ventspils", "liepaja", "liep%C4%81ja", "daugavpils", "jelgava", "jurmala", "j%C5%ABrmala"},
	"romania":      PresetLocations{"romania", "bucharest", "cluj", "iasi", "timisoara", "craiova", "brasov", "sibiu", "constanta", "oradea", "galati", "ploesti", "pitesti", "arad", "bacau"},
	"belarus":      PresetLocations{"belarus", "minsk", "brest,belarus", "grodno", "gomel", "vitebsk", "mogilev", "slutsk", "borisov", "pinsk", "baranovichi", "bobruisk", "soligorsk"},
	"malta":        PresetLocations{"malta", "birgu", "bormla", "mdina", "qormi", "rabat", "senglea", "si%C4%A1%C4%A1iewi", "valletta", "zabbar", "zebbu%C4%A1", "zejtun"},
	"rwanda":       PresetLocations{"rwanda", "kigali", "butare", "muhanga", "ruhengeri", "gisenyi"},
	"saudi arabia": PresetLocations{"Saudi", "KSA", "Riyadh", "Mecca"},
	"morocco":      PresetLocations{"morocco", "casablanca", "fez", "tangier", "marrakesh", "sal%E9", "meknes", "rabat", "oujda", "kenitra", "agadir", "tetouan", "temara", "safi", "mohammedia", "khouribga", "el%2Bjadida"},
	"uzbekistan":   PresetLocations{"uzbekistan", "tashkent", "namangan", "samarkand", "andijan", "nukus", "bukhara", "qarshi", "fergana"},
	"malaysia":     PresetLocations{"malaysia", "kuala%2Blumpur", "kajang", "klang", "subang", "penang", "ipoh", "selangor", "melaka", "johor", "sabah"},
	"afghanistan":  PresetLocations{"afghanistan", "kabul", "kandahar", "herat", "mazar-e-sharif", "jalalabad", "ghazni", "nangarhar"},
	"venezuela":    PresetLocations{"venezuela", "caracas", "maracaibo", "barquisimeto", "guayana", "matur%C3%ADn", "zulia", "bolivar"},
	"ghana":        PresetLocations{"ghana", "accra", "kumasi", "sekondi", "ashaiman", "sunyani", "tamale", "tema"},
	"angola":       PresetLocations{"angola", "luanda", "huambo", "lobito", "benguela"},
	"nepal":        PresetLocations{"nepal", "kathmandu", "pokhara", "lalitpur", "bharatpur", "birgunj", "biratnagar", "janakpur", "ghorahi"},
	"yemen":        PresetLocations{"yemen", "sana%27a", "taiz", "aden", "mukalla", "ibb", "marib"},
	"mozambique":   PresetLocations{"mozambique", "maputo", "matola", "nampula", "beira", "sofala", "chimoio", "tete", "quelimane"},
	"ivory coast":  PresetLocations{"ivory", "abidjan", "bouak%C3%A9", "daloa", "yamoussoukro"}}

func Preset(name string) []string {
	return PRESETS[name]
}
