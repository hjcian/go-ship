package tagfetcher

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestImageTags_Len(t *testing.T) {
	tags := ImageTags{
		{Name: "tag1"},
		{Name: "tag2"},
		{Name: "tag3"},
	}
	if tags.Len() != 3 {
		t.Errorf("expected length 3, got %d", tags.Len())
	}

	emptyTags := ImageTags{}
	if emptyTags.Len() != 0 {
		t.Errorf("expected length 0, got %d", emptyTags.Len())
	}
}

func TestImageTags_GetLatestPushed(t *testing.T) {
	now := time.Now()
	t1 := now.Add(-2 * time.Hour)
	t2 := now.Add(-1 * time.Hour)
	t3 := now

	tags := ImageTags{
		{Name: "tag1", TagLastPushed: t1, Digest: "digest1"},
		{Name: "tag2", TagLastPushed: t2, Digest: "digest2"},
		{Name: "tag3", TagLastPushed: t3, Digest: "digest3"},
	}

	latest := tags.GetLatestPushed()
	if latest == nil {
		t.Fatal("expected non-nil latest tag")
	}
	if latest.Name != "tag3" {
		t.Errorf("expected latest tag to be 'tag3', got '%s'", latest.Name)
	}
	if !latest.TagLastPushed.Equal(t3) {
		t.Errorf("expected latest TagLastPushed to be %v, got %v", t3, latest.TagLastPushed)
	}

	// Test with empty slice
	var emptyTags ImageTags
	if emptyTags.GetLatestPushed() != nil {
		t.Error("expected nil for empty tags slice")
	}

	// Test with single element
	singleTag := ImageTags{
		{Name: "only", TagLastPushed: now, Digest: "digest"},
	}
	latest = singleTag.GetLatestPushed()
	assert.NotNil(t, latest)
	assert.Equal(t, "only", latest.Name)
	assert.Equal(t, now, latest.TagLastPushed)

	// Test the _test_data, unmarshal it and check the latest tag
	var testTags ImageTags
	if err := json.Unmarshal([]byte(_test_data), &testTags); err != nil {
		t.Fatalf("failed to unmarshal test data: %v", err)
	}
	latest = testTags.GetLatestPushed()
	assert.NotNil(t, latest)
	assert.Equal(t, "latest", latest.Name)
	assert.Equal(t, "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea", latest.Digest)
	assert.Equal(t, "2025-07-22T22:05:59Z", latest.TagLastPushed.Format(time.RFC3339))
}

var _test_data = `
[
  {
    "name": "latest",
    "tag_last_pushed": "2025-07-22T22:05:59.257041Z",
    "digest": "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea"
  },
  {
    "name": "bookworm",
    "tag_last_pushed": "2025-07-22T22:05:53.294706Z",
    "digest": "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea"
  },
  {
    "name": "8.0.3-bookworm",
    "tag_last_pushed": "2025-07-22T22:05:41.933364Z",
    "digest": "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea"
  },
  {
    "name": "8.0.3",
    "tag_last_pushed": "2025-07-22T22:05:37.148926Z",
    "digest": "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea"
  },
  {
    "name": "8.0-bookworm",
    "tag_last_pushed": "2025-07-22T22:05:34.77781Z",
    "digest": "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea"
  },
  {
    "name": "8.0",
    "tag_last_pushed": "2025-07-22T22:05:30.06208Z",
    "digest": "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea"
  },
  {
    "name": "8-bookworm",
    "tag_last_pushed": "2025-07-22T22:05:27.904972Z",
    "digest": "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea"
  },
  {
    "name": "8",
    "tag_last_pushed": "2025-07-22T22:05:22.243954Z",
    "digest": "sha256:f957ce918b51f3ac10414244bedd0043c47db44a819f98b9902af1bd9d0afcea"
  },
  {
    "name": "7.4.5-bookworm",
    "tag_last_pushed": "2025-07-22T22:05:16.556168Z",
    "digest": "sha256:49061c0de5717822cf6702ef3197f6817f26b98d46765c308c1e23e6b261997b"
  },
  {
    "name": "7.4.5",
    "tag_last_pushed": "2025-07-22T22:05:12.008417Z",
    "digest": "sha256:49061c0de5717822cf6702ef3197f6817f26b98d46765c308c1e23e6b261997b"
  },
  {
    "name": "7.4-bookworm",
    "tag_last_pushed": "2025-07-22T22:05:09.713676Z",
    "digest": "sha256:49061c0de5717822cf6702ef3197f6817f26b98d46765c308c1e23e6b261997b"
  },
  {
    "name": "7.4",
    "tag_last_pushed": "2025-07-22T22:05:05.079445Z",
    "digest": "sha256:49061c0de5717822cf6702ef3197f6817f26b98d46765c308c1e23e6b261997b"
  },
  {
    "name": "7.2.10-bookworm",
    "tag_last_pushed": "2025-07-22T22:05:02.872977Z",
    "digest": "sha256:c5dbd5bb88464705a302a3a1114d610e4baa13179c7c67946047528790cc19dd"
  },
  {
    "name": "7.2.10",
    "tag_last_pushed": "2025-07-22T22:04:57.900375Z",
    "digest": "sha256:c5dbd5bb88464705a302a3a1114d610e4baa13179c7c67946047528790cc19dd"
  },
  {
    "name": "7.2-bookworm",
    "tag_last_pushed": "2025-07-22T22:04:55.680737Z",
    "digest": "sha256:c5dbd5bb88464705a302a3a1114d610e4baa13179c7c67946047528790cc19dd"
  },
  {
    "name": "7.2",
    "tag_last_pushed": "2025-07-22T22:04:50.318094Z",
    "digest": "sha256:c5dbd5bb88464705a302a3a1114d610e4baa13179c7c67946047528790cc19dd"
  },
  {
    "name": "7-bookworm",
    "tag_last_pushed": "2025-07-22T22:04:44.518289Z",
    "digest": "sha256:49061c0de5717822cf6702ef3197f6817f26b98d46765c308c1e23e6b261997b"
  },
  {
    "name": "7",
    "tag_last_pushed": "2025-07-22T22:04:39.250044Z",
    "digest": "sha256:49061c0de5717822cf6702ef3197f6817f26b98d46765c308c1e23e6b261997b"
  },
  {
    "name": "6.2.19-bookworm",
    "tag_last_pushed": "2025-07-22T22:04:33.559918Z",
    "digest": "sha256:37440b1f783e3dcabb49eaaaa1bb52a619275072eadf9dadbaa271e5b69bfe2c"
  },
  {
    "name": "6.2.19",
    "tag_last_pushed": "2025-07-22T22:04:28.94437Z",
    "digest": "sha256:37440b1f783e3dcabb49eaaaa1bb52a619275072eadf9dadbaa271e5b69bfe2c"
  },
  {
    "name": "6.2-bookworm",
    "tag_last_pushed": "2025-07-22T22:04:26.74805Z",
    "digest": "sha256:37440b1f783e3dcabb49eaaaa1bb52a619275072eadf9dadbaa271e5b69bfe2c"
  },
  {
    "name": "6.2",
    "tag_last_pushed": "2025-07-22T22:04:21.761783Z",
    "digest": "sha256:37440b1f783e3dcabb49eaaaa1bb52a619275072eadf9dadbaa271e5b69bfe2c"
  },
  {
    "name": "6-bookworm",
    "tag_last_pushed": "2025-07-22T22:04:19.247059Z",
    "digest": "sha256:37440b1f783e3dcabb49eaaaa1bb52a619275072eadf9dadbaa271e5b69bfe2c"
  },
  {
    "name": "6",
    "tag_last_pushed": "2025-07-22T22:04:13.008242Z",
    "digest": "sha256:37440b1f783e3dcabb49eaaaa1bb52a619275072eadf9dadbaa271e5b69bfe2c"
  },
  {
    "name": "8.2-rc1-bookworm",
    "tag_last_pushed": "2025-07-22T19:05:18.686552Z",
    "digest": "sha256:2991fb723c00b3bf3eb17b552fde827caad4f307f6b6218ec80e8972670198bf"
  },
  {
    "name": "8.2-rc1",
    "tag_last_pushed": "2025-07-22T19:05:13.540993Z",
    "digest": "sha256:2991fb723c00b3bf3eb17b552fde827caad4f307f6b6218ec80e8972670198bf"
  },
  {
    "name": "7.4.5-alpine3.21",
    "tag_last_pushed": "2025-07-17T19:05:22.540929Z",
    "digest": "sha256:bb186d083732f669da90be8b0f975a37812b15e913465bb14d845db72a4e3e08"
  },
  {
    "name": "7.4.5-alpine",
    "tag_last_pushed": "2025-07-17T19:05:20.318346Z",
    "digest": "sha256:bb186d083732f669da90be8b0f975a37812b15e913465bb14d845db72a4e3e08"
  },
  {
    "name": "7.4-alpine3.21",
    "tag_last_pushed": "2025-07-17T19:05:15.663156Z",
    "digest": "sha256:bb186d083732f669da90be8b0f975a37812b15e913465bb14d845db72a4e3e08"
  },
  {
    "name": "7.4-alpine",
    "tag_last_pushed": "2025-07-17T19:05:13.297594Z",
    "digest": "sha256:bb186d083732f669da90be8b0f975a37812b15e913465bb14d845db72a4e3e08"
  }
]
`
