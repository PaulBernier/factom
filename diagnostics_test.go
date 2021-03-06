// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factom_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/FactomProject/factom"

	"testing"
)

func TestUnmarshalDiagnostics(t *testing.T) {
	js := []byte(`{"name":"FNode0","id":"38bab1455b7bd7e5efd15c53c777c79d0c988e9210f1da49a99d95b3a6417be9","publickey":"cc1985cdfae4e32b5a454dfda8ce5e1361558482684f3367649c3ad852c8e31a","role":"Follower","leaderheight":186663,"currentminute":2,"currentminuteduration":189834229037,"previousminuteduration":1554309715022145853,"balancehash":"66634eb6aa5816f5b39786647d2083ff6e274daadfe2223cf7cec1546bdf3f43","tempbalancehash":"66632cfbbd63a61ef57dbdbea9c5161cbecccb56469ab439ada14761a55bfad5","lastblockfromdbstate":false,"syncing":{"status":"Syncing EOMs","received":24,"expected":26,"missing":["8888880180b0290bbb670e399af48e57a227c939a2d20f6b0e147d24f995a6ef","8888887d00cb1f7a13a94f5fc07cf77ee2b0b5be460c0c2915a838b9458baea7"]},"authset":{"leaders":[{"id":"8888880180b0290bbb670e399af48e57a227c939a2d20f6b0e147d24f995a6ef","vm":14,"listheight":3,"listlength":4,"nextnil":0},{"id":"88888807e4f3bbb9a2b229645ab6d2f184224190f83e78761674c2362aca4425","vm":15,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888880a1ad522921100a0fdbc42ab4e701cae15d71c5ce414ec74ecd2b6d201","vm":16,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888880c67754f737fd59310d7dbf2df08daed3b161347fbe76ca24517373911","vm":17,"listheight":3,"listlength":4,"nextnil":0},{"id":"88888820713a0bf32f29bfdc4b48ecd7aab8942651ccfbacf5c31ad70c16b3aa","vm":18,"listheight":3,"listlength":4,"nextnil":0},{"id":"88888824716497c80f5bd509cd59049cd957dc4ff64ddd5ad77505a758a2c2dc","vm":19,"listheight":3,"listlength":4,"nextnil":0},{"id":"88888828d6ed2ceb15b47ab83c1ca300b5c54f485685825ae5e988d2e21f767a","vm":20,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888882aa2a521373b5b4129f958383d5dd6708644346cef43e07c5c22e6bfdb","vm":21,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888882adec124fa6afa8094e873516e2ca7d18aa740d55da95f0004b6c222e4","vm":22,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888883a40c004ba51834dd2599f271b30e3251180295f099886754d1b993667","vm":23,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888885a035a14a64da6df04f6c459587a05a72e8796b428b067a04ff97e5680","vm":24,"listheight":3,"listlength":4,"nextnil":0},{"id":"888888631d6561b8d7ae4cc2b55dd39228219a00ac930d80e3d5bd06c2d8cadc","vm":25,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888886a135d497b86e6ffc6a1352d42dab1765074d14d9d2da97dbc747f1a60","vm":0,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888886ff14cef50365b785eb3cefab5bc30175d022be06ed412391a82645376","vm":1,"listheight":18,"listlength":18,"nextnil":0},{"id":"8888887529d62b6d3d702bafb06f11ef825ec2fd54c978c1e1809a7eedba1514","vm":2,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888887d00cb1f7a13a94f5fc07cf77ee2b0b5be460c0c2915a838b9458baea7","vm":3,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888887e062010156e9b7a838b30f11719003746373ff95ceb745ff9d79075b2","vm":4,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888888fa04184f2ee054205ca4a542c5045c3e0e391b639f317b06112cf2f7f","vm":5,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888889b229588d21762ab8df380bacef09a8ed0774eed76e6c7f0d0b9d5ee4d","vm":6,"listheight":3,"listlength":4,"nextnil":0},{"id":"8888889dcbea07cbac9814511d1e1b8b4560e75b06132a75aa2b066a507ae755","vm":7,"listheight":3,"listlength":4,"nextnil":0},{"id":"888888aac90d34ba84830046be9bdbc0c12d39045ce8f3f6c95f0beca4636629","vm":8,"listheight":3,"listlength":4,"nextnil":0},{"id":"888888aecf9b47fd6f12015aea0260ad5eb179b6678b5d1f96a3c5c9131ce5e8","vm":9,"listheight":3,"listlength":4,"nextnil":0},{"id":"888888b4a7105ec297396592b8e9448b504a8fb41b82ee26e23068ff0e4549d0","vm":10,"listheight":3,"listlength":4,"nextnil":0},{"id":"888888e3eded221899e7be0816cef4c61e1911d674f019d591cf2ecb6250d608","vm":11,"listheight":3,"listlength":4,"nextnil":0},{"id":"888888f00138ff5aeb903a6ccc98d5b0e1b0944edebf34328a68154854abc6ed","vm":12,"listheight":18,"listlength":18,"nextnil":0},{"id":"888888f4d59308deaa587498e5e1c4e0228a190eba50c9ad23b604da1cbd8c77","vm":13,"listheight":3,"listlength":4,"nextnil":0}],"audits":[{"id":"88888804081f44513658c1565558f7e2dfb9b3b992763d88349e635db4b83101","online":true},{"id":"8888880834dfe56c8c5827026eec58a8a71e3496fb76035c8710f7b708a3daf6","online":true},{"id":"8888880852d183e020b7fbf764adcb6559c6b9c53f2851446f2506fddf387015","online":true},{"id":"88888808e14a4e802ba540c86c2d6345d69a4938e77dd7beecbc0da6dc47b3e3","online":true},{"id":"8888881f8fd587a9e4b5163112b7332df01e7c5b11294652106a9b8e4e87ec24","online":true},{"id":"88888821e559c76aca86e03582abdb07e87095bfa239a0b5feb3943e6c85b0a3","online":true},{"id":"88888852a0821dd227735c50c5016741de67786432b9644c5d695b5ce7e42e58","online":true},{"id":"888888609b7bb56a8f184c624b12f9db9cb356f9eb20fc4153aee3479d1b6cd7","online":true},{"id":"88888862fd0d52de4d5f10e4d80956f9ff8b63f98564ca7673d08237d0e4b4de","online":true},{"id":"88888863888a959a8f9b6a664244f9ab00d1ab0cb2708cc0977643b40ca9ead4","online":true},{"id":"888888655866a003faabd999c7b0a7c908af17d63fd2ac2951dc99e1ad2a14f4","online":true},{"id":"8888887f5125bfc597a05eca2db64298b88a9233dafdeb44bc0db7d55ee035aa","online":true},{"id":"8888888e08f7ece967b74e92d54f65f5e874b5ec6fa2ae507895a699539462ac","online":true},{"id":"8888889987d4ca5f0687ef2392327ec6c29bc71ef981e10e7473fdd63a296cba","online":true},{"id":"8888889e4fbbcc0032e6a2ce517d39fc90cce1189a46d7cebfff4b8bc230744c","online":true},{"id":"888888b02c99d60b41357f5e28a1e96f4b90d370f6a460f2062cdda64e3fc7a4","online":true},{"id":"888888cfc9326b0daad96ab68b0fb39f094db4b44ac35f06bfccddf440e95b45","online":true},{"id":"888888d2a9a5a37b19dc8093e058945591f6d6ec5a6bbcf50727c6c86d98c50c","online":true},{"id":"888888d7bb4d1c5c667ca663fccfae26bdea6198604582aae038d962c5e937ab","online":true},{"id":"888888d92dfbd1a69353e2be28740eb09a7299d82b0f72a1885f1be663187dc9","online":true},{"id":"888888dda15d7ad44c3286d66cc4f82e6fc07ed88de4d13ac9a182199593cac1","online":true},{"id":"888888e419097489fae57af195f1d9dbaf7741ed0004a6109588214f7a97c5cb","online":true},{"id":"888888e61ebcf0fa694c07840de5618b87cc7f4b1a13cf23512b1c6ccf0cb17d","online":true},{"id":"888888f5b2bb6c049fb0790d908a8e3f0ecee5166616ac06c5444cc48180abaf","online":true},{"id":"888888ff0fa60c17e33b6173068c7dcacdc4d0ea55df6fbdd0ff2ae2db13917f","online":true}]},"elections":{"inprogress":true,"vmindex":1,"fedindex":13,"fedid":"8888886ff14cef50365b785eb3cefab5bc30175d022be06ed412391a82645376","round":1}}`)

	d := new(Diagnostics)
	err := json.Unmarshal(js, d)
	if err != nil {
		t.Error(err)
	}
	t.Log(d)
}

// a local factomd api server must be running for this test to pass!
func TestGetDiagnostics(t *testing.T) {
	factomdResponse := `{
	   "jsonrpc": "2.0",
	   "id": 1,
	   "result": {
	      "name": "FNode0",
	      "id": "38bab1455b7bd7e5efd15c53c777c79d0c988e9210f1da49a99d95b3a6417be9",
	      "publickey": "cc1985cdfae4e32b5a454dfda8ce5e1361558482684f3367649c3ad852c8e31a",
	      "role": "Follower",
	      "leaderheight": 192669,
	      "currentminute": 0,
	      "currentminuteduration": 2335868444,
	      "previousminuteduration": 1557938081229465300,
	      "balancehash": "266960096c4e0e016a9dff266f25c91039eed9c28c8f7339e29ec724b60aaafe",
	      "tempbalancehash": "26695dbc339e7316aea2683faf839c1b7b1ee2313db792112588118df066aa35",
	      "lastblockfromdbstate": false,
	      "syncing": {
	         "status": "Processing"
	      },
	      "authset": {
	         "leaders": [
	            {
	               "id": "8888880180b0290bbb670e399af48e57a227c939a2d20f6b0e147d24f995a6ef",
	               "vm": 10,
	               "listheight": 0,
	               "listlength": 0,
	               "nextnil": 0
	            }, {
	               "id": "88888807e4f3bbb9a2b229645ab6d2f184224190f83e78761674c2362aca4425",
	               "vm": 11,
	               "listheight": 0,
	               "listlength": 0,
	               "nextnil": 0
	            }, {
	               "id": "8888880a1ad522921100a0fdbc42ab4e701cae15d71c5ce414ec74ecd2b6d201",
	               "vm": 12,
	               "listheight": 0,
	               "listlength": 0,
	               "nextnil": 0
	            }, {
	               "id": "8888880c67754f737fd59310d7dbf2df08daed3b161347fbe76ca24517373911",
	               "vm": 13,
	               "listheight": 0,
	               "listlength": 0,
	               "nextnil": 0
	            }, {
	               "id": "888888f4d59308deaa587498e5e1c4e0228a190eba50c9ad23b604da1cbd8c77",
	               "vm": 9,
	               "listheight": 0,
	               "listlength": 0,
	               "nextnil": 0
	            }
	         ],
	         "audits": [
	            {
	               "id": "88888804081f44513658c1565558f7e2dfb9b3b992763d88349e635db4b83101",
	               "online": false
	            }, {
	               "id": "8888880834dfe56c8c5827026eec58a8a71e3496fb76035c8710f7b708a3daf6",
	               "online": false
	            }, {
	               "id": "88888808e14a4e802ba540c86c2d6345d69a4938e77dd7beecbc0da6dc47b3e3",
	               "online": false
	            }, {
	               "id": "8888881f8fd587a9e4b5163112b7332df01e7c5b11294652106a9b8e4e87ec24",
	               "online": false
	            }, {
	               "id": "88888821e559c76aca86e03582abdb07e87095bfa239a0b5feb3943e6c85b0a3",
	               "online": false
	            }, {
	               "id": "88888824716497c80f5bd509cd59049cd957dc4ff64ddd5ad77505a758a2c2dc",
	               "online": false
	            }, {
	               "id": "88888852a0821dd227735c50c5016741de67786432b9644c5d695b5ce7e42e58",
	               "online": false
	            }, {
	               "id": "888888ff0fa60c17e33b6173068c7dcacdc4d0ea55df6fbdd0ff2ae2db13917f",
	               "online": false
	            }
	         ]
	      },
	      "elections": {
	         "inprogress": false
	      }
	   }
	}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, factomdResponse)
	}))
	defer ts.Close()

	SetFactomdServer(ts.URL[7:])

	d, err := GetDiagnostics()
	if err != nil {
		t.Error(err)
	}
	t.Log(d)
}
