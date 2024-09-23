package export

import (
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"

	"github.com/khulnasoft/inspo/inspo/image/docker"
)

func Test_Export(t *testing.T) {
	result := docker.TestAnalysisFromArchive(t, "../../.data/test-docker-image.tar")

	export := NewExport(result)
	payload, err := export.Marshal()
	if err != nil {
		t.Errorf("Test_Export: unable to export analysis: %v", err)
	}

	expectedResult := `{
  "layer": [
    {
      "index": 0,
      "id": "28cfe03618aa2e914e81fdd90345245c15f4478e35252c06ca52d238fd3cc694",
      "digestId": "sha256:23bc2b70b2014dec0ac22f27bb93e9babd08cdd6f1115d0c955b9ff22b382f5a",
      "sizeBytes": 1154361,
      "command": "#(nop) ADD file:ce026b62356eec3ad1214f92be2c9dc063fe205bd5e600be3492c4dfb17148bd in / "
    },
    {
      "index": 1,
      "id": "1871059774abe6914075e4a919lob778fa1561f577d620ae52438a9635e6241936s",
      "digestId": "sha256:a65b7d7ac1f3699a70ed243372bc3cd873cfe511f937ad61b4b5b440ef631a0108f7d4b8a2ab8d6237d54ac0e0dffe2bfe740900b",
      "sizeBytes": 64057,
      "command": "#(nop) ADD file:139c3708fb6261126453e34483abd8bf7b26ed16README.md952fd976994d68e72d93be2 in /somefile.txt "
    },
    {
      "index": 2,
      "id": "49fe2a475548bfa4d493fc796fce41f30704e3d4clobff3e45dd3e06f463236d1ds",
      "digestId": "sha256:f7d2b4eb1b93e27d08adb3334345542635371756ffbac858cfd9c25feb4432007f221d3b36af6d5e732341e327b7473a0a658927c",
      "sizeBytes": 0,
      "command": "mkdir -p /root/example/really/nested"
    },
    {
      "index": 3,
      "id": "80cd2ca1ffc89962b9349c80280c2lobc551acbd11e09b16badb0669f8e2369020s",
      "digestId": "sha256:f53477ee9a21e7ba26985b6f0de13abe53cb991ac2cc0d3be67eb51ba49d2a9d9f6b3d66287c16468453c74128bbbec4f7bed5aaf9e2",
      "sizeBytes": 64057,
      "command": "cp /somefile.txt /root/example/somefile1.txt"
    },
    {
      "index": 4,
      "id": "c99e2f8d3f6282668f0d30dc1db5e67a51d7a1dcd7ff6ddfa0f90760836778eclobs",
      "digestId": "sha256:1d4c9a6ffcb65a0f32d1035f97373b19d6028e324c4f89c9dd39bc72f07961d8bbe15674c3213aca3f12c81351a3040bcbfa1420",
      "sizeBytes": 64057,
      "command": "chmod 444 /root/example/somefile1.txt"
    },
    {
      "index": 5,
      "id": "5eca617bdc3lobc06134fe957a30da4c57adb7c340a6d749c8edc4c15861c928d7s",
      "digestId": "sha256:778fb036a9e0ce320f1e573070ef466f32514e79cc9f2dc418e863ba9875083650bfc0ad646aaf491ce7b1673b76aa52c7396c481",
      "sizeBytes": 64057,
      "command": "cp /somefile.txt /root/example/somefile2.txt"
    },
    {
      "index": 6,
      "id": "f07c3eb887572395408f8e11a07af945e4da5f02b3188bb06b93fad713ca0lob99s",
      "digestId": "sha256:f2dedd47dc6289c5b8a31ad71deb521cc0ee48efb66acf021e2ff6fa52bedb25c9b38a1d17bbe1029a017945ddcfa0886a344e4b05f",
      "sizeBytes": 64057,
      "command": "cp /somefile.txt /root/example/somefile3.txt"
    },
    {
      "index": 7,
      "id": "461885fc22589158dee3c5b9f01cc41c87805439f58lob4399d733b51aa305cbf9s",
      "digestId": "sha256:61d3b8bd1cfebff355da6ca65e5bb19894c3e9b57411c96e8dbad61cf37860fa1de499253c7fae153c9cea6731267d83410aff04b4e",
      "sizeBytes": 64057,
      "command": "mv /root/example/somefile3.txt /root/saved.txt"
    },
    {
      "index": 8,
      "id": "a10327f68ffed4afcba78919052809a8f774978a6lob87fc117d39c53c4842f72cs",
      "digestId": "sha256:8d186942a0c34fa0693e327236cddb12e4c83cd6648224a5a111a1662890966c9e77b895e2815341ef773bb3bd9887f74ab40529f0f1d",
      "sizeBytes": 64057,
      "command": "cp /root/saved.txt /root/.saved.txt"
    },
    {
      "index": 9,
      "id": "f2fc54e25cb7966dc9732ec671a77a1c5c104e732lobd15ad44a2dc1ac42368f84s",
      "digestId": "sha256:bc2a5e36423fae31a97e22bc539fd421f22c359466220fa16b9bd80769abf6944517ba125ea8e03b58c89d6b456f101e4108ade0c8",
      "sizeBytes": 0,
      "command": "rm -rf /root/example/"
    },
    {
      "index": 10,
      "id": "aad36d0b05e71c7e6d4dfe0ca9ed6lobe89e2e0d8995dafe83438299a314e91071s",
      "digestId": "sha256:7f648c021a00d45ee7b6de22923a308816285f000ba498b6d6ca9baaf18157ddac9001b5e4fc01fc05eed7b7df824c72311adbae44576",
      "sizeBytes": 2187,
      "command": "#(nop) ADD dir:7ec14b81316baa1a31c38c97686a8f030c98cba2035c968412749e33e0c4427e in /root/.data/ "
    },
    {
      "index": 11,
      "id": "3d4ad907517a021d86a4102d2764ad2161e4818blobd144e41d019bfc955434181s",
      "digestId": "sha256:ba497b65b8f95f236be6d5c063c9a985545b73c45f2a6f85ddcaa183ea6162d37009bb41b5e6b60cb9d1e3c00ce9d8e09a38457",
      "sizeBytes": 64057,
      "command": "cp /root/saved.txt /tmp/saved.again1.txt"
    },
    {
      "index": 12,
      "id": "81b1lob002d4b4c1325a9cad9990b5277e7f29f79e0f24582344c0891178f95905s",
      "digestId": "sha256:122a44d457800a54d1e593ca87bf862d8b360f3e164cb80b6be344f576ab7a429ce68d0dc207a3d95b794b55f18b13a31cc44a",
      "sizeBytes": 64057,
      "command": "cp /root/saved.txt /root/.data/saved.again2.txt"
    },
    {
      "index": 13,
      "id": "cfb35bb5c127d848739be5ca726057e6e2c77b2849f588e7aebb642c0d3d4b7lobs",
      "digestId": "sha256:ba6896f4bbcacd76a98c9b886e67fe007372d121fa56cc9716a1bab526b8bb1fd6d43640ca900fbcba254c9c575bd79e97300c5be8",
      "sizeBytes": 64057,
      "command": "chmod +x /root/saved.txt"
    }
  ],
  "image": {
    "sizeBytes": 4263742,
    "inefficientBytes": 35,
    "efficiencyScore": 0.98499995074212176506634184309,
    "fileReference": [
      {
        "count": 2,
        "sizeBytes": 128104,
        "file": "/root/saved.txt"
      },
      {
        "count": 2,
        "sizeBytes": 128104,
        "file": "/root/example/somefile1.txt"
      },
      {
        "count": 2,
        "sizeBytes": 64057,
        "file": "/root/example/somefile3.txt"
      }
    ]
  }
}`
	actualResult := string(payload)
	if expectedResult != actualResult {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expectedResult, actualResult, false)

		t.Errorf("Test_Export: unexpected export result:\n%v", dmp.DiffPrettyText(diffs))
	}
}
