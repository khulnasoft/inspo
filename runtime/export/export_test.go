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
      "id": "blobs",
      "digestId": "sha256:58f32e9504c8eb248292326a1975174b0560f7a3ad1b75880b9571c4eb7144a0",
      "sizeBytes": 4269678,
      "command": "BusyBox 1.37.0 (glibc), Debian 12"
    },
    {
      "index": 1,
      "id": "blobs",
      "digestId": "sha256:85cdb07e71d49a058b08493884eb1f27c9bab1f828d0e4609e28e967e419bcba",
      "sizeBytes": 1628,
      "command": "ADD README.md /somefile.txt # buildkit"
    },
    {
      "index": 2,
      "id": "blobs",
      "digestId": "sha256:daca5b56c35419d1f90bbbe2f248b57fc8fe23aa1544de09f66d535c4b04c696",
      "sizeBytes": 0,
      "command": "RUN /bin/sh -c mkdir -p /root/example/really/nested # buildkit"
    },
    {
      "index": 3,
      "id": "blobs",
      "digestId": "sha256:8f16d72c66faa2c2a6e858c0b3f388bc0f3485c82438f0acd38353a7f24edcf3",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c cp /somefile.txt /root/example/somefile1.txt # buildkit"
    },
    {
      "index": 4,
      "id": "blobs",
      "digestId": "sha256:bdd91c0accd5f2f5ca425a34f0b1b9f4fc4240de38162da4cdb2add428eb5ddd",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c chmod 444 /root/example/somefile1.txt # buildkit"
    },
    {
      "index": 5,
      "id": "blobs",
      "digestId": "sha256:2c5734409001054090704d689c6c87ae578e488e02911475afa7b2ee45d0ba98",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c cp /somefile.txt /root/example/somefile2.txt # buildkit"
    },
    {
      "index": 6,
      "id": "blobs",
      "digestId": "sha256:536042a303aa7337c9b669da5ef1007a8ed9382bfb6cae9ee589ea60400943b7",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c cp /somefile.txt /root/example/somefile3.txt # buildkit"
    },
    {
      "index": 7,
      "id": "blobs",
      "digestId": "sha256:587e756b742eb0fa9dfdae31e085ff1debe3e954f0f98495f5c18cb6afd082e7",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c mv /root/example/somefile3.txt /root/saved.txt # buildkit"
    },
    {
      "index": 8,
      "id": "blobs",
      "digestId": "sha256:3065bf55b675855573b902ec00530a41f75ceb2d347cc3f97e315a84280f4401",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c cp /root/saved.txt /root/.saved.txt # buildkit"
    },
    {
      "index": 9,
      "id": "blobs",
      "digestId": "sha256:9902797c49c049a9f90f7552bc06f2ae777869648e9ba642e5398b6706e5726e",
      "sizeBytes": 0,
      "command": "RUN /bin/sh -c rm -rf /root/example/ # buildkit"
    },
    {
      "index": 10,
      "id": "blobs",
      "digestId": "sha256:546bbe15caebfdd936f0679f15f2a666fdd3d29a4ed2639eb5667ce654eb4f48",
      "sizeBytes": 2186,
      "command": "ADD .scripts/ /root/.data/ # buildkit"
    },
    {
      "index": 11,
      "id": "blobs",
      "digestId": "sha256:1343732693e6c8ea9ac54df638a159faef67592463b1841b43020131b8eabe77",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c cp /root/saved.txt /tmp/saved.again1.txt # buildkit"
    },
    {
      "index": 12,
      "id": "blobs",
      "digestId": "sha256:40b5fb651fc7e657186a93d2b0d097cdf47c5634439613ebbeeae6413504db88",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c cp /root/saved.txt /root/.data/saved.again2.txt # buildkit"
    },
    {
      "index": 13,
      "id": "blobs",
      "digestId": "sha256:9b084d4d974d2c9ee7bd720398af6537fbf21caebd400bde33403dce5e2a67fc",
      "sizeBytes": 1628,
      "command": "RUN /bin/sh -c chmod +x /root/saved.txt # buildkit"
    }
  ],
  "image": {
    "sizeBytes": 4288144,
    "inefficientBytes": 8140,
    "efficiencyScore": 0.9988619098662441,
    "fileReference": [
      {
        "count": 2,
        "sizeBytes": 3256,
        "file": "/root/saved.txt"
      },
      {
        "count": 2,
        "sizeBytes": 3256,
        "file": "/root/example/somefile1.txt"
      },
      {
        "count": 2,
        "sizeBytes": 1628,
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
