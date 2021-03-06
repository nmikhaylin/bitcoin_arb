package api_calls

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "time"
  "bytes"
  "encoding/json"
)

var kTestAlgoInfo = []byte(`{"Success":true,"Message":"1,000 api calls remaining.","Data":[{"CoinName":"Digitalcoin-X11","CoinTag":"DGC","Algorithm":"X11","Difficulty":16.59818099,"BlockReward":5,"BlockCount":2561291,"ProfitRatio":3566.5830028783921,"AvgProfitRatio":2918.7030337265246,"Exchange":"Cryptopia","ExchangeRate":1.002E-05,"ExchangeVolume":76268.1770073,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":40,"HealthStatus":"Healthy"},{"CoinName":"Spots","CoinTag":"SPT","Algorithm":"Scrypt","Difficulty":0.11714456,"BlockReward":11.25,"BlockCount":975987,"ProfitRatio":2320.7844030641536,"AvgProfitRatio":192.20176260907954,"Exchange":"Cryptopia","ExchangeRate":9E-08,"ExchangeVolume":3942.74481653,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":70,"HealthStatus":"Healthy"},{"CoinName":"Dash","CoinTag":"DASH","Algorithm":"X11","Difficulty":103037.25163735,"BlockReward":1.80147602,"BlockCount":666055,"ProfitRatio":1327.2406018194411,"AvgProfitRatio":5839.82495972526,"Exchange":"Poloniex","ExchangeRate":0.06494019,"ExchangeVolume":144420.20193209,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":150,"HealthStatus":"Healthy"},{"CoinName":"Maxcoin","CoinTag":"MAX","Algorithm":"Keccak","Difficulty":30.71375742,"BlockReward":16,"BlockCount":2003559,"ProfitRatio":1261.6284693621881,"AvgProfitRatio":1231.2737266640054,"Exchange":"Bter","ExchangeRate":1E-06,"ExchangeVolume":0,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Bitbar","CoinTag":"BTB","Algorithm":"Scrypt","Difficulty":92.82731986,"BlockReward":0.11748814316721246,"BlockCount":219991,"ProfitRatio":1201.5686502242045,"AvgProfitRatio":2878.8139614000561,"Exchange":"Cryptopia","ExchangeRate":0.00378705,"ExchangeVolume":39.37956622,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":600,"HealthStatus":"Healthy"},{"CoinName":"MonetaryUnit","CoinTag":"MUE","Algorithm":"Quark","Difficulty":356524.62595057,"BlockReward":40,"BlockCount":2230832,"ProfitRatio":1181.1081269585125,"AvgProfitRatio":5922.8379620293726,"Exchange":"Bittrex","ExchangeRate":3.5E-05,"ExchangeVolume":1332650.51432025,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":40,"HealthStatus":"Healthy"},{"CoinName":"Casinocoin","CoinTag":"CSC","Algorithm":"Scrypt","Difficulty":1.87291018,"BlockReward":1,"BlockCount":1960243,"ProfitRatio":947.46664763720969,"AvgProfitRatio":1219.752773947195,"Exchange":"Cryptopia","ExchangeRate":7.34E-06,"ExchangeVolume":209379.46946886,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":30,"HealthStatus":"Healthy"},{"CoinName":"ArtByte","CoinTag":"ABY","Algorithm":"Scrypt","Difficulty":40.08036638,"BlockReward":400,"BlockCount":791614,"ProfitRatio":881.63596083317191,"AvgProfitRatio":2076.5702678180419,"Exchange":"Bittrex","ExchangeRate":3.7E-07,"ExchangeVolume":10549842.6198164,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":120,"HealthStatus":"Healthy"},{"CoinName":"CHNCoin","CoinTag":"CNC","Algorithm":"Scrypt","Difficulty":34.76499331,"BlockReward":88,"BlockCount":1896947,"ProfitRatio":867.81740774313857,"AvgProfitRatio":32733.404852851334,"Exchange":"Bter","ExchangeRate":1.44E-06,"ExchangeVolume":166625.579,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Ethereum","CoinTag":"ETH","Algorithm":"EtHash","Difficulty":357115156213669,"BlockReward":5,"BlockCount":3668919,"ProfitRatio":855.08430569895461,"AvgProfitRatio":4558.6908318896367,"Exchange":"Poloniex","ExchangeRate":0.05652565,"ExchangeVolume":610794.33902156,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":15,"HealthStatus":"Healthy"},{"CoinName":"Ethereum-Classic","CoinTag":"ETC","Algorithm":"EtHash","Difficulty":29261020233478,"BlockReward":5,"BlockCount":3649966,"ProfitRatio":846.50775918635054,"AvgProfitRatio":6340.0414230775041,"Exchange":"Bter","ExchangeRate":0.00459061,"ExchangeVolume":1695.272,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":15,"HealthStatus":"Healthy"},{"CoinName":"Goldcoin","CoinTag":"GLD","Algorithm":"Scrypt","Difficulty":32.20856811,"BlockReward":4,"BlockCount":619798,"ProfitRatio":824.47892337635551,"AvgProfitRatio":2371.3709600372395,"Exchange":"Bittrex","ExchangeRate":2.815E-05,"ExchangeVolume":49071.3445298,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":120,"HealthStatus":"Healthy"},{"CoinName":"StartCoin","CoinTag":"START","Algorithm":"X11","Difficulty":97.89447887,"BlockReward":10,"BlockCount":1243642,"ProfitRatio":761.75447723614957,"AvgProfitRatio":1201.7899535345055,"Exchange":"Bittrex","ExchangeRate":6.46E-06,"ExchangeVolume":1372535.01848772,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"GameCredits","CoinTag":"GAME","Algorithm":"Scrypt","Difficulty":5167.22802855,"BlockReward":25,"BlockCount":1659809,"ProfitRatio":736.46499929622826,"AvgProfitRatio":2557.4233792763262,"Exchange":"Poloniex","ExchangeRate":0.00066,"ExchangeVolume":1153554.99953003,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":90,"HealthStatus":"Healthy"},{"CoinName":"Novacoin","CoinTag":"NVC","Algorithm":"Scrypt","Difficulty":6525.79801927,"BlockReward":5.7831936832734021,"BlockCount":336611,"ProfitRatio":717.59885827787423,"AvgProfitRatio":2919.127647077637,"Exchange":"BTCe","ExchangeRate":0.00353,"ExchangeVolume":18624.54305,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":450,"HealthStatus":"Healthy"},{"CoinName":"Florincoin","CoinTag":"FLO","Algorithm":"Scrypt","Difficulty":113.79379108,"BlockReward":25,"BlockCount":2154183,"ProfitRatio":702.31835346871969,"AvgProfitRatio":2383.2720280939361,"Exchange":"Bittrex","ExchangeRate":1.4E-05,"ExchangeVolume":742387.30857277,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":40,"HealthStatus":"Healthy"},{"CoinName":"BATA","CoinTag":"BTA","Algorithm":"Scrypt","Difficulty":15.95606816,"BlockReward":0.390625,"BlockCount":652773,"ProfitRatio":661.63635464714389,"AvgProfitRatio":1613.5937291740415,"Exchange":"Bittrex","ExchangeRate":0.00011992,"ExchangeVolume":50541.58166224,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":90,"HealthStatus":"Healthy"},{"CoinName":"Quark","CoinTag":"QRK","Algorithm":"Quark","Difficulty":703.84813969,"BlockReward":1,"BlockCount":4276229,"ProfitRatio":643.08061345493309,"AvgProfitRatio":5434.2734294822485,"Exchange":"Cryptopia","ExchangeRate":1.52E-06,"ExchangeVolume":4173.33863289,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":30,"HealthStatus":"Healthy"},{"CoinName":"Belacoin","CoinTag":"BELA","Algorithm":"Scrypt","Difficulty":1564.94991681,"BlockReward":50,"BlockCount":682862,"ProfitRatio":628.31637810502252,"AvgProfitRatio":1917.2733112417468,"Exchange":"Poloniex","ExchangeRate":8.83E-05,"ExchangeVolume":2272486.57823161,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":120,"HealthStatus":"Healthy"},{"CoinName":"Zclassic","CoinTag":"ZCL","Algorithm":"Equihash","Difficulty":148067.57406466,"BlockReward":12.5,"BlockCount":103882,"ProfitRatio":618.31935065416894,"AvgProfitRatio":3683.2462087347885,"Exchange":"Bittrex","ExchangeRate":0.00531355,"ExchangeVolume":38483.77141904,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":150,"HealthStatus":"Healthy"},{"CoinName":"Zcash","CoinTag":"ZEC","Algorithm":"Equihash","Difficulty":1471774.95173052,"BlockReward":10.00024776,"BlockCount":110232,"ProfitRatio":599.53455870186542,"AvgProfitRatio":4330.9162311239106,"Exchange":"Poloniex","ExchangeRate":0.06432753,"ExchangeVolume":37852.44035886,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":150,"HealthStatus":"Healthy"},{"CoinName":"Worldcoin","CoinTag":"WDC","Algorithm":"Scrypt","Difficulty":73.04556814,"BlockReward":16.31422503,"BlockCount":3335934,"ProfitRatio":595.75055700636369,"AvgProfitRatio":1563.7342874229103,"Exchange":"Cryptopia","ExchangeRate":1.213E-05,"ExchangeVolume":90213.51260622,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":30,"HealthStatus":"Healthy"},{"CoinName":"Digitalcoin-Scrypt","CoinTag":"DGC","Algorithm":"Scrypt","Difficulty":18.73409184,"BlockReward":5,"BlockCount":2561290,"ProfitRatio":585.60995006893029,"AvgProfitRatio":1003.8630166720903,"Exchange":"Cryptopia","ExchangeRate":1.002E-05,"ExchangeVolume":76268.1770073,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":40,"HealthStatus":"Healthy"},{"CoinName":"DNotes","CoinTag":"NOTE","Algorithm":"Scrypt","Difficulty":105.9184949,"BlockReward":21.43437499,"BlockCount":1687050,"ProfitRatio":577.67008545416684,"AvgProfitRatio":2380.9500108025372,"Exchange":"Poloniex","ExchangeRate":1.308E-05,"ExchangeVolume":5055551.57070843,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"DigiByte","CoinTag":"DGB","Algorithm":"Scrypt","Difficulty":340.60052378,"BlockReward":891.17417,"BlockCount":4449372,"ProfitRatio":576.41647918786157,"AvgProfitRatio":2459.0418931765707,"Exchange":"Poloniex","ExchangeRate":1.01E-06,"ExchangeVolume":1789410140.4193959,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":15,"HealthStatus":"Healthy"},{"CoinName":"Litecoin","CoinTag":"LTC","Algorithm":"Scrypt","Difficulty":179053.70456602,"BlockReward":25,"BlockCount":1199933,"ProfitRatio":555.43422524669359,"AvgProfitRatio":2390.167682591652,"Exchange":"BTCe","ExchangeRate":0.01841006,"ExchangeVolume":258950.90229,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":150,"HealthStatus":"Healthy"},{"CoinName":"Einsteinium","CoinTag":"EMC2","Algorithm":"Scrypt","Difficulty":26.80051246,"BlockReward":8,"BlockCount":1411001,"ProfitRatio":502.39489254562636,"AvgProfitRatio":1689.4357335184061,"Exchange":"Poloniex","ExchangeRate":8E-06,"ExchangeVolume":33305301.40749136,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"WildBeastBitcoin","CoinTag":"WBB","Algorithm":"Scrypt","Difficulty":118.92480436,"BlockReward":0.625,"BlockCount":794783,"ProfitRatio":482.09003456506883,"AvgProfitRatio":918.71652701603944,"Exchange":"Bittrex","ExchangeRate":0.0004411,"ExchangeVolume":1326.71854547,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Monero","CoinTag":"XMR","Algorithm":"CryptoNight","Difficulty":8790168829,"BlockReward":7.7426183753770044,"BlockCount":1305311,"ProfitRatio":465.79302159805979,"AvgProfitRatio":2416.5551942822476,"Exchange":"Poloniex","ExchangeRate":0.02075007,"ExchangeVolume":637843.91929227,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":120,"HealthStatus":"Healthy"},{"CoinName":"Quatloo","CoinTag":"QTL","Algorithm":"Scrypt","Difficulty":26.05308151,"BlockReward":4.2525,"BlockCount":528496,"ProfitRatio":463.47436354686033,"AvgProfitRatio":1783.2701065245978,"Exchange":"Bittrex","ExchangeRate":1.381E-05,"ExchangeVolume":465425.3487371,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":1335,"HealthStatus":"Healthy"},{"CoinName":"Verge","CoinTag":"XVG","Algorithm":"Scrypt","Difficulty":28.23519908,"BlockReward":1560,"BlockCount":1176185,"ProfitRatio":450.64790342887608,"AvgProfitRatio":1339.8479383584981,"Exchange":"Bittrex","ExchangeRate":4E-08,"ExchangeVolume":177068274.809462,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":30,"HealthStatus":"Healthy"},{"CoinName":"Peercoin","CoinTag":"PPC","Algorithm":"SHA-256","Difficulty":2101954922.6964602,"BlockReward":46.698262150113507,"BlockCount":300244,"ProfitRatio":389.98016562779111,"AvgProfitRatio":296.19275067666166,"Exchange":"BTCe","ExchangeRate":0.0016,"ExchangeVolume":259538.14186,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":600,"HealthStatus":"Healthy"},{"CoinName":"ScoreCoin","CoinTag":"SCORE","Algorithm":"Scrypt","Difficulty":21.18292332,"BlockReward":250,"BlockCount":61212,"ProfitRatio":357.13062416405779,"AvgProfitRatio":486.81714968118348,"Exchange":"Cryptopia","ExchangeRate":1.6E-07,"ExchangeVolume":175896.94437727,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Megacoin","CoinTag":"MEC","Algorithm":"Scrypt","Difficulty":212.3692705,"BlockReward":12.5,"BlockCount":800796,"ProfitRatio":355.73472142425959,"AvgProfitRatio":1399.315377784839,"Exchange":"Cryptopia","ExchangeRate":3.2E-05,"ExchangeVolume":56465.92506158,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":150,"HealthStatus":"Healthy"},{"CoinName":"eMark","CoinTag":"DEM","Algorithm":"SHA-256","Difficulty":8297322.88819126,"BlockReward":50,"BlockCount":1241637,"ProfitRatio":341.94849412151689,"AvgProfitRatio":1150.8634067672692,"Exchange":"Cryptopia","ExchangeRate":5.58E-06,"ExchangeVolume":626970.89949714,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":120,"HealthStatus":"Healthy"},{"CoinName":"Gulden","CoinTag":"NLG","Algorithm":"Scrypt","Difficulty":1382.0763477,"BlockReward":100,"BlockCount":521031,"ProfitRatio":335.08862599685705,"AvgProfitRatio":1712.3196165984059,"Exchange":"Bittrex","ExchangeRate":2.505E-05,"ExchangeVolume":2865098.04367721,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":150,"HealthStatus":"Healthy"},{"CoinName":"Digitalcoin-SHA-256","CoinTag":"DGC","Algorithm":"SHA-256","Difficulty":1578112.78723725,"BlockReward":5,"BlockCount":2561291,"ProfitRatio":294.9846133546348,"AvgProfitRatio":-1137.6081109451391,"Exchange":"Cryptopia","ExchangeRate":1.002E-05,"ExchangeVolume":76268.1770073,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":40,"HealthStatus":"Healthy"},{"CoinName":"Sexcoin","CoinTag":"SXC","Algorithm":"Scrypt","Difficulty":3.40190764,"BlockReward":12.5,"BlockCount":2214968,"ProfitRatio":288.86508149511673,"AvgProfitRatio":1748.3623377828594,"Exchange":"Cryptopia","ExchangeRate":4.5E-07,"ExchangeVolume":1452039.51491259,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Nyancoin","CoinTag":"NYAN","Algorithm":"Scrypt","Difficulty":1.33594944,"BlockReward":42.125,"BlockCount":1671743,"ProfitRatio":266.51297062261983,"AvgProfitRatio":2288.0938554494728,"Exchange":"Cryptopia","ExchangeRate":5E-08,"ExchangeVolume":926297.29218132,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Zetacoin","CoinTag":"ZET","Algorithm":"SHA-256","Difficulty":73594.50573176,"BlockReward":1,"BlockCount":6754212,"ProfitRatio":235.0645077617184,"AvgProfitRatio":-571.03836297021167,"Exchange":"Bleutrade","ExchangeRate":2.16E-06,"ExchangeVolume":16450.46357802,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":30,"HealthStatus":"Healthy"},{"CoinName":"Emerald","CoinTag":"EMD","Algorithm":"Scrypt","Difficulty":0.20774688,"BlockReward":0.5,"BlockCount":5327602,"ProfitRatio":234.98297306004332,"AvgProfitRatio":1252.6852332830765,"Exchange":"Cryptopia","ExchangeRate":6.1E-07,"ExchangeVolume":153715.35786443,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":20,"HealthStatus":"Healthy"},{"CoinName":"Auroracoin","CoinTag":"AUR","Algorithm":"Scrypt","Difficulty":346.70303782,"BlockReward":2.5,"BlockCount":795416,"ProfitRatio":196.34502528051439,"AvgProfitRatio":2064.1325988625904,"Exchange":"Bittrex","ExchangeRate":0.00018517,"ExchangeVolume":60220.99163983,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":300,"HealthStatus":"Healthy"},{"CoinName":"Blakecoin","CoinTag":"BLC","Algorithm":"Blake-256","Difficulty":4447.70258221,"BlockReward":25.00814842,"BlockCount":569609,"ProfitRatio":167.83925104056365,"AvgProfitRatio":1691.6179355248098,"Exchange":"Cryptopia","ExchangeRate":1.7E-06,"ExchangeVolume":637386.74798947,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":180,"HealthStatus":"Healthy"},{"CoinName":"Vertcoin","CoinTag":"VTC","Algorithm":"Lyra2REv2","Difficulty":3238.57556723,"BlockReward":50,"BlockCount":713998,"ProfitRatio":139.7106794964424,"AvgProfitRatio":175.32217655619436,"Exchange":"Bittrex","ExchangeRate":0.00020816,"ExchangeVolume":1326481.82881227,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":150,"HealthStatus":"Healthy"},{"CoinName":"Groestlcoin","CoinTag":"GRS","Algorithm":"Groestl","Difficulty":40.03541318,"BlockReward":6.06041577,"BlockCount":1579917,"ProfitRatio":131.95549596348752,"AvgProfitRatio":826.61734053816679,"Exchange":"Bittrex","ExchangeRate":1.612E-05,"ExchangeVolume":159617.4788242,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Bitcoin","CoinTag":"BTC","Algorithm":"SHA-256","Difficulty":521974519553.628,"BlockReward":12.5,"BlockCount":465346,"ProfitRatio":100,"AvgProfitRatio":100,"Exchange":"Bitfinex","ExchangeRate":1643.8,"ExchangeVolume":25768.39793461,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":600,"HealthStatus":"Healthy"},{"CoinName":"Diamond","CoinTag":"DMD","Algorithm":"Groestl","Difficulty":42.71957588,"BlockReward":0.2,"BlockCount":2152843,"ProfitRatio":89.825757215342151,"AvgProfitRatio":973.662290120285,"Exchange":"Bittrex","ExchangeRate":0.00044555,"ExchangeVolume":3334.75704203,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":100,"HealthStatus":"Healthy"},{"CoinName":"Orbitcoin","CoinTag":"ORB","Algorithm":"NeoScrypt","Difficulty":0.46677351,"BlockReward":1,"BlockCount":2161855,"ProfitRatio":25.347049117758125,"AvgProfitRatio":145.58259724835233,"Exchange":"Bleutrade","ExchangeRate":4.507E-05,"ExchangeVolume":192.98357276,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":180,"HealthStatus":"Healthy"},{"CoinName":"Feathercoin","CoinTag":"FTC","Algorithm":"NeoScrypt","Difficulty":20.90329044,"BlockReward":80,"BlockCount":1705738,"ProfitRatio":7.3464077656830273,"AvgProfitRatio":993.85845330703262,"Exchange":"Bittrex","ExchangeRate":2.078E-05,"ExchangeVolume":1829585.61577816,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Aricoin","CoinTag":"ARI","Algorithm":"Scrypt","Difficulty":3.13638513,"BlockReward":200,"BlockCount":485097,"ProfitRatio":-6.4327721087555023,"AvgProfitRatio":167.79637787296952,"Exchange":"Cryptopia","ExchangeRate":1E-08,"ExchangeVolume":14373.89473823,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":180,"HealthStatus":"Healthy"},{"CoinName":"Bottlecaps","CoinTag":"CAP","Algorithm":"Scrypt","Difficulty":3.76647063,"BlockReward":10,"BlockCount":2794498,"ProfitRatio":-6.5709967311904594,"AvgProfitRatio":1739.1751458026413,"Exchange":"Cryptopia","ExchangeRate":2.4E-07,"ExchangeVolume":258182.26110993,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"CannabisCoin","CoinTag":"CANN","Algorithm":"X11","Difficulty":0.0455908,"BlockReward":0.00020027,"BlockCount":2155833,"ProfitRatio":-8.1283037365594843,"AvgProfitRatio":-0.15045041068659307,"Exchange":"Bittrex","ExchangeRate":2.85E-06,"ExchangeVolume":528312.8411689,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":42,"HealthStatus":"Healthy"},{"CoinName":"Securecoin","CoinTag":"SRC","Algorithm":"Quark","Difficulty":99306.38303776,"BlockReward":5,"BlockCount":1551142,"ProfitRatio":-11.01324732556983,"AvgProfitRatio":53.306448305012466,"Exchange":"Bter","ExchangeRate":2.2E-07,"ExchangeVolume":0,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Phoenixcoin","CoinTag":"PXC","Algorithm":"NeoScrypt","Difficulty":0.11571242,"BlockReward":25,"BlockCount":1370253,"ProfitRatio":-33.328816597282149,"AvgProfitRatio":99.585844188685158,"Exchange":"Cryptopia","ExchangeRate":1.9E-07,"ExchangeVolume":210721.95034435,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":90,"HealthStatus":"Healthy"},{"CoinName":"Mazacoin","CoinTag":"MZC","Algorithm":"SHA-256","Difficulty":7049715.96186275,"BlockReward":1000,"BlockCount":806044,"ProfitRatio":-108.60515489845017,"AvgProfitRatio":-244.16796330193918,"Exchange":"Bittrex","ExchangeRate":1.1E-07,"ExchangeVolume":13411305.7223952,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":120,"HealthStatus":"Healthy"},{"CoinName":"Unbreakable","CoinTag":"UNB","Algorithm":"SHA-256","Difficulty":85888520.5720191,"BlockReward":50,"BlockCount":67202,"ProfitRatio":-133.53864822229252,"AvgProfitRatio":-612.1887505695853,"Exchange":"Bittrex","ExchangeRate":2.509E-05,"ExchangeVolume":79893.30052216,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":300,"HealthStatus":"Healthy"},{"CoinName":"Dogecoin","CoinTag":"DOGE","Algorithm":"Scrypt","Difficulty":70925.63388655,"BlockReward":10000,"BlockCount":1705514,"ProfitRatio":-156.14844118874834,"AvgProfitRatio":-788.44358858261569,"Exchange":"Bter","ExchangeRate":8.7E-07,"ExchangeVolume":45415229.977,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Viacoin","CoinTag":"VIA","Algorithm":"Scrypt","Difficulty":8337.55900972,"BlockReward":0.625,"BlockCount":3659445,"ProfitRatio":-188.64469623976782,"AvgProfitRatio":-912.77635855843778,"Exchange":"Bittrex","ExchangeRate":0.00014512,"ExchangeVolume":109652.95448272,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":24,"HealthStatus":"Healthy"},{"CoinName":"Joulecoin","CoinTag":"XJO","Algorithm":"SHA-256","Difficulty":195274.95456173,"BlockReward":8,"BlockCount":2581058,"ProfitRatio":-232.40271919617226,"AvgProfitRatio":-241.8822986730413,"Exchange":"Cryptopia","ExchangeRate":2.6E-07,"ExchangeVolume":70478.40153001,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":45,"HealthStatus":"Healthy"},{"CoinName":"Tigercoin","CoinTag":"TGC","Algorithm":"SHA-256","Difficulty":26317.76756498,"BlockReward":1,"BlockCount":2272347,"ProfitRatio":-365.70618398413728,"AvgProfitRatio":-1853.8879690072652,"Exchange":"Cryptopia","ExchangeRate":1.4E-07,"ExchangeVolume":13878.46682899,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":45,"HealthStatus":"Healthy"},{"CoinName":"Acoin","CoinTag":"ACOIN","Algorithm":"SHA-256","Difficulty":18314.67749226,"BlockReward":0.01,"BlockCount":3116193,"ProfitRatio":-374.47902585614321,"AvgProfitRatio":-2225.2531772612533,"Exchange":"Cryptopia","ExchangeRate":9.1E-06,"ExchangeVolume":373843.68807753,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":40,"HealthStatus":"Healthy"},{"CoinName":"Crown","CoinTag":"CRW","Algorithm":"SHA-256","Difficulty":411067531.25440496,"BlockReward":4.5,"BlockCount":1366407,"ProfitRatio":-384.5181892991443,"AvgProfitRatio":-558.15601927370585,"Exchange":"Bittrex","ExchangeRate":0.0004172,"ExchangeVolume":66105.53171564,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"Terracoin","CoinTag":"TRC","Algorithm":"SHA-256","Difficulty":611691991.85346,"BlockReward":20,"BlockCount":987456,"ProfitRatio":-488.94665402917565,"AvgProfitRatio":-2479.4736643732613,"Exchange":"Cryptopia","ExchangeRate":1.193E-05,"ExchangeVolume":152426.94416573,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":120,"HealthStatus":"Healthy"},{"CoinName":"TEKcoin","CoinTag":"TEK","Algorithm":"SHA-256","Difficulty":54410.61328942,"BlockReward":1,"BlockCount":2278875,"ProfitRatio":-489.50891983605612,"AvgProfitRatio":-2492.7443158524957,"Exchange":"Cryptopia","ExchangeRate":2E-08,"ExchangeVolume":682579.99913279,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":60,"HealthStatus":"Healthy"},{"CoinName":"HAMRadioCoin","CoinTag":"HAM","Algorithm":"SHA-256","Difficulty":80323.57882925,"BlockReward":0.390625,"BlockCount":352291,"ProfitRatio":-493.83523505213714,"AvgProfitRatio":-3768.1068104477581,"Exchange":"Bleutrade","ExchangeRate":4E-08,"ExchangeVolume":34110.38486261,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":180,"HealthStatus":"Healthy"},{"CoinName":"Namecoin","CoinTag":"NMC","Algorithm":"SHA-256","Difficulty":264459529140.41,"BlockReward":25,"BlockCount":338305,"ProfitRatio":-494.83438172695287,"AvgProfitRatio":-2552.2073946308192,"Exchange":"Bter","ExchangeRate":0.001635,"ExchangeVolume":106.589,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":600,"HealthStatus":"Healthy"},{"CoinName":"Unobtanium","CoinTag":"UNO","Algorithm":"SHA-256","Difficulty":217574660.80586,"BlockReward":0.00195312,"BlockCount":946091,"ProfitRatio":-496.48700855392758,"AvgProfitRatio":-2554.207138798356,"Exchange":"Cryptopia","ExchangeRate":0.00985393,"ExchangeVolume":715.39909633,"IsBlockExplorerOnline":true,"IsExchangeOnline":true,"Message":"","BlockTimeInSeconds":180,"HealthStatus":"Healthy"},{"CoinName":"Curecoin","CoinTag":"CURE","Algorithm":"SHA-256","Difficulty":15034387.8269431,"BlockReward":13,"BlockCount":170524,"ProfitRatio":358.65073503365795,"AvgProfitRatio":-462.06318084566487,"Exchange":"Poloniex","ExchangeRate":3.966E-05,"ExchangeVolume":52987.37016497,"IsBlockExplorerOnline":true,"IsExchangeOnline":false,"Message":"Warning: One or more exchange is offline which may cause out dated estimated expected cryptocurrency earnings to be returned.","BlockTimeInSeconds":600,"HealthStatus":"Unhealthy"},{"CoinName":"UniversalCurrency","CoinTag":"UNIT","Algorithm":"SHA-256","Difficulty":385128.21924564,"BlockReward":5,"BlockCount":2071082,"ProfitRatio":-86.491030793388688,"AvgProfitRatio":NaN,"Exchange":"Bittrex","ExchangeRate":1.27E-06,"ExchangeVolume":2929931.51047181,"IsBlockExplorerOnline":true,"IsExchangeOnline":false,"Message":"Warning: One or more exchange is offline which may cause out dated estimated expected cryptocurrency earnings to be returned.","BlockTimeInSeconds":60,"HealthStatus":"Unhealthy"},{"CoinName":"Myriadcoin-Scrypt","CoinTag":"MYR","Algorithm":"Scrypt","Difficulty":70735.23106822,"BlockReward":250,"BlockCount":2049367,"ProfitRatio":-191.59133942960102,"AvgProfitRatio":-1264.28547275137,"Exchange":"Poloniex","ExchangeRate":2.1E-07,"ExchangeVolume":19781005.930943329,"IsBlockExplorerOnline":true,"IsExchangeOnline":false,"Message":"Warning: One or more exchange is offline which may cause out dated estimated expected cryptocurrency earnings to be returned.Warning: Re-syncing the blockchain","BlockTimeInSeconds":300,"HealthStatus":"Unhealthy"},{"CoinName":"Myriadcoin-SHA-256","CoinTag":"MYR","Algorithm":"SHA-256","Difficulty":307674386.19219,"BlockReward":250,"BlockCount":2049367,"ProfitRatio":-494.43251417153374,"AvgProfitRatio":-3517.2224686586546,"Exchange":"Poloniex","ExchangeRate":2.1E-07,"ExchangeVolume":19781005.930943329,"IsBlockExplorerOnline":true,"IsExchangeOnline":false,"Message":"Warning: One or more exchange is offline which may cause out dated estimated expected cryptocurrency earnings to be returned.Warning: Re-syncing the blockchain","BlockTimeInSeconds":300,"HealthStatus":"Unhealthy"}]}`)

var myClient = &http.Client{Timeout: 10 * time.Second}

type AlgoInfo struct {
  Id int
  UnitMultiplier float32
}

var kTradedAlgoInfoMap = map[string]AlgoInfo{
  "Scrypt": {0, 1e6},
  "SHA-256": {1, 1e5},
  "X11": {3, 1e3},
  "EtHash": {20, 1e3}, // DaggerHashimoto
}

var kCoinScaling = map[string]float32 {
  "Novacoin": 36584613.011,
  "DigiByte": 36584613.011,
  //  "Gulden": 31227042.000000,
  // "Belacoin": 32117838.000000,
  "GameCredits": 36584613.011,
  "Ethereum": 156362570005334000.00,
  "Dash": 36584613.011,
  "DNotes": 36584613.011,
  "Litecoin": 36584613.011,
  // "Bitcoin": 22972243968.000000,
  // "Crown": 83304002879488.000000,
  // "Viacoin": 7561455104.000000,
  // "Namecoin": 8950542193655808.000000,
}

var CurrentBitcoinPrice float32 = 0.0

const kApiKey string = "20aeeaae2f8341d2842ef67af9ab44dd"
const kApiKeyInfoEndpoint string = "http://www.coinwarz.com/v1/api/apikeyinfo"
const kProfitabilityEndpoint string = "http://www.coinwarz.com/v1/api/profitability/"
const kNiceHashEndPoint string = "https://www.nicehash.com/api"

const kSupportedAlgos string = "all"
const kMinVolumeInBTC float32 = 20.0

type ApiKeyInfo struct {
  Success bool `json:"Success"`
  Message string `json:"Message"`
  Data struct{
    ApiKey string `json:"ApiKey"`
    ApiUsageLimit int `json:"ApiUsageLimit"`
    ApiUsageRemaining int `json:"ApiUsageRemaining"`
    DailyUsageLimit int `json:"DailyUsageLimit"`
    DailyUsageRemaining int `json:"DailyUsageRemaining"`
  } `json:"Data"`
}

func GetApiKeyInfo() (*ApiKeyInfo, error) { 
  req, err := http.NewRequest("GET", kApiKeyInfoEndpoint, nil)
  if (err != nil) {
    return nil, err
  }
  q := req.URL.Query()
  q.Add("apikey", kApiKey)
  req.URL.RawQuery = q.Encode()
  resp, err := myClient.Get(req.URL.String())
  if (err != nil) {
    return nil, err
  }
  body, err := ioutil.ReadAll(resp.Body)
  var s = new(ApiKeyInfo)
  err = json.Unmarshal(body[:], &s)
  if (err != nil) {
    return nil, err
  }
  return s, nil
}

type CoinData struct {
    CoinName string `json:"CoinName"`
    CoinTag string `json:"CoinTag"`
    Algorithm string `json:"Algorithm"`
    HealthStatus string `json:"HealthStatus"`
    Difficulty float32 `json:"Difficulty"`
    BlockReward float32 `json:"BlockReward"`
    BlockCount int `json:"BlockCount"`
    ProfitRatio float32 `json:"ProfitRatio"`
    ExchangeVolume float32 `json:"ExchangeVolume"`
    ExchangeRate float32 `json:"ExchangeRate"`
    AvgProfitRatio float32 `json:"AvgProfitRatio"`
}

type MiningProfitability struct {
  Success bool `json:"Success"`
  Message string `json:"Message"`
  Data []CoinData `json:"Data"`
}

type MarketRate struct {
  Algo string
  Rate float32
}

func GetProfitabilityInfo(rates map[string]MarketRate) ([]CoinData, error) {
  body := bytes.Replace(kTestAlgoInfo, []byte("NaN"), []byte("null"), -1)
  var err error
  if (true) {
    req, err := http.NewRequest("GET", kProfitabilityEndpoint, nil)
    if (err != nil) {
      return nil, err
    }
    q := req.URL.Query()
    q.Add("apikey", kApiKey)
    // q.Add("algo", "all")
    q.Add("algo", kSupportedAlgos)
    q.Add("sha256HashRate", fmt.Sprintf(
        "%f", 1.0 / rates["SHA-256"].Rate *
        kTradedAlgoInfoMap["SHA-256"].UnitMultiplier))
    q.Add("sha256PowerCost", "0.0")
    q.Add("scryptHashRate", fmt.Sprintf(
        "%f", 1.0 / rates["Scrypt"].Rate *
        kTradedAlgoInfoMap["Scrypt"].UnitMultiplier))
    q.Add("scryptPowerCost", "0.0")
    q.Add("x11HashRate", fmt.Sprintf(
        "%f", 1.0 / rates["X11"].Rate *
        kTradedAlgoInfoMap["X11"].UnitMultiplier))
    q.Add("x11PowerCost", "0.0")
    q.Add("ethashHashRate", fmt.Sprintf(
        "%f", 1.0 / rates["EtHash"].Rate *
        kTradedAlgoInfoMap["EtHash"].UnitMultiplier))
    q.Add("ethashPowerCost", "0.0")
    req.URL.RawQuery = q.Encode()
    resp, err := myClient.Get(req.URL.String())
    if (err != nil) {
      return nil, err
    }
    body, err = ioutil.ReadAll(resp.Body)
    body = bytes.Replace(body, []byte("NaN"), []byte("null"), -1)
  }
  var profit = new(MiningProfitability)
  err = json.Unmarshal(body[:], &profit)
  if (err != nil) {
    return nil, err
  }
  GetCurrentBitcoinPrice(profit)
  FilterUnHealthy(profit)
  FilterLowVolume(profit)
  FilterAlgorithms(profit)
  FilterSupportedCoins(profit)
  CalculateProfitability(profit, rates)
  return profit.Data, nil
}

func GetCurrentBitcoinPrice(profit *MiningProfitability) {
  for _, v := range profit.Data {
    if v.CoinName == "Bitcoin" {
      CurrentBitcoinPrice = v.ExchangeRate
      fmt.Printf("CurrentBitcoinPrice: %f\n", CurrentBitcoinPrice)
      break
    }
  }
}

func FilterLowVolume(profit *MiningProfitability) {
  new_data := []CoinData{}
  for _, v := range profit.Data {
    if ((v.ExchangeVolume * v.ExchangeRate) > kMinVolumeInBTC) {
      new_data = append(new_data, v)
    }
  }
  profit.Data = new_data
}

func FilterUnHealthy(profit *MiningProfitability) {
  new_data := []CoinData{}
  for _, v := range profit.Data {
    if v.HealthStatus == "Healthy" {
      new_data = append(new_data, v)
    }
  }
  profit.Data = new_data
}


func FilterAlgorithms(profit *MiningProfitability) {
  new_data := []CoinData{}
  for _, v := range profit.Data {
    if _, pres := kTradedAlgoInfoMap[v.Algorithm] ; pres {
      new_data = append(new_data, v)
    }
  }
  profit.Data = new_data
}

func FilterSupportedCoins(profit *MiningProfitability) {
  new_data := []CoinData{}
  for _, v := range profit.Data {
    if _, pres := kCoinScaling[v.CoinName] ; pres {
      new_data = append(new_data, v)
    }
  }
  profit.Data = new_data
}

// Assumes all coins have scaling.
func CalculateProfitability(profit *MiningProfitability, rates map[string]MarketRate) {
  new_data := []CoinData{}
  for _, v := range profit.Data {
    if scale, pres := kCoinScaling[v.CoinName] ; pres {
      v.ProfitRatio = 1000 * (scale * v.BlockReward * v.ExchangeRate /
          rates[v.Algorithm].Rate / CurrentBitcoinPrice / v.Difficulty)
      new_data = append(new_data, v)
    }
  }
  profit.Data = new_data
}

type AlgoMarket struct {
  Algo string
  Location int
  Orders []NiceHashOrder
}

type NiceHashMarket struct {
  AlgoMarkets []AlgoMarket
}

type NiceHashOrder struct {
  LimitSpeed float32 `json:"limit_speed,string"`
  Alive bool `json:"alive"`
  Price float32 `json:"price,string"`
  OrderId int `json:"id"`
  OrderType int `json:"type"`
  Algo int `json:"algo"`
  Workers int `json:"workers"`
  AcceptedSpeed float32 `json:"accepted_speed,string"`
}

type NiceHashMarketInfo struct {
  Result struct {
    Orders []NiceHashOrder `json:"orders"`
  } `json:"result"`
}

func GetNiceHashMarket() (*NiceHashMarket, error) {
  ret := new(NiceHashMarket)
  algo_markets := []AlgoMarket{}
  for name, info := range kTradedAlgoInfoMap {
    req, err := http.NewRequest("GET", kNiceHashEndPoint, nil)
    if (err != nil) {
      return nil, err
    }
    q := req.URL.Query()
    q.Add("method", "orders.get")
    q.Add("algo", fmt.Sprintf("%d", info.Id))
    // Only US currently.
    q.Add("location", "1")
    req.URL.RawQuery = q.Encode()
    resp, err := myClient.Get(req.URL.String())
    if (err != nil) {
      return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    var info = new(NiceHashMarketInfo)
    err = json.Unmarshal(body[:], &info)
    if (err != nil) {
      return nil, err
    }
    algo_markets = append(
      algo_markets, AlgoMarket{Algo:name, Location:1,
                               Orders:info.Result.Orders[:]})
  }
  ret.AlgoMarkets = algo_markets
  return ret, nil
}
