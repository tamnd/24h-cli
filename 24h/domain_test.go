package h24

import (
	"testing"

	"github.com/tamnd/any-cli/kit"
)

func TestDomainInfo(t *testing.T) {
	info := Domain{}.Info()
	if info.Scheme != "h24" {
		t.Errorf("Scheme = %q, want h24", info.Scheme)
	}
	if len(info.Hosts) == 0 || info.Hosts[0] != Host {
		t.Errorf("Hosts = %v, want [%s]", info.Hosts, Host)
	}
	if info.Identity.Binary != "h24" {
		t.Errorf("Binary = %q, want h24", info.Identity.Binary)
	}
}

func TestClassify(t *testing.T) {
	cases := []struct {
		in      string
		wantTyp string
		wantID  string
		wantErr bool
	}{
		{"https://www.24h.com.vn/bong-da/bai-viet-c123-190001.html", "article", "bong-da/bai-viet-c123-190001.html", false},
		{"bong-da/tin-c45-1234567.html", "article", "bong-da/tin-c45-1234567.html", false},
		{"bong-da", "category", "bong-da", false},
		{"xa-hoi", "category", "xa-hoi", false},
		{"", "", "", true},
		{"not-a-category-or-url", "", "", true},
	}
	for _, tc := range cases {
		typ, id, err := Domain{}.Classify(tc.in)
		if tc.wantErr {
			if err == nil {
				t.Errorf("Classify(%q): want error", tc.in)
			}
			continue
		}
		if err != nil {
			t.Errorf("Classify(%q): %v", tc.in, err)
			continue
		}
		if typ != tc.wantTyp || id != tc.wantID {
			t.Errorf("Classify(%q) = (%q,%q), want (%q,%q)", tc.in, typ, id, tc.wantTyp, tc.wantID)
		}
	}
}

func TestLocate(t *testing.T) {
	cases := []struct {
		typ, id, want string
		wantErr       bool
	}{
		{"article", "bong-da/bai-viet-c123-190001.html", "https://www.24h.com.vn/bong-da/bai-viet-c123-190001.html", false},
		{"category", "bong-da", "https://www.24h.com.vn/bong-da/", false},
		{"unknown", "x", "", true},
	}
	for _, tc := range cases {
		got, err := Domain{}.Locate(tc.typ, tc.id)
		if tc.wantErr {
			if err == nil {
				t.Errorf("Locate(%q,%q): want error", tc.typ, tc.id)
			}
			continue
		}
		if err != nil || got != tc.want {
			t.Errorf("Locate(%q,%q) = (%q,%v), want (%q,nil)", tc.typ, tc.id, got, err, tc.want)
		}
	}
}

func TestHostWiring(t *testing.T) {
	h, err := kit.Open()
	if err != nil {
		t.Fatal(err)
	}

	a := &Article{
		ID:       "bong-da/bai-viet-c123-190001.html",
		URL:      "https://www.24h.com.vn/bong-da/bai-viet-c123-190001.html",
		Category: "bong-da",
	}
	u, err := h.Mint(a)
	if err != nil {
		t.Fatalf("Mint: %v", err)
	}
	if want := "h24://article/bong-da/bai-viet-c123-190001.html"; u.String() != want {
		t.Errorf("Mint = %q, want %q", u.String(), want)
	}
}
