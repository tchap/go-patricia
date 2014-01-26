// Copyright (c) 2014 The go-patricia AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package patricia

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

const (
	success = true
	failure = false
)

type testData struct {
	key    string
	value  interface{}
	retVal bool
}

// Tests -----------------------------------------------------------------------

func TestTrie_InsertDifferentPrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepaneeeeeeeeeeeeee", "Pepan Zdepan", success},
		{"Honzooooooooooooooo", "Honza Novak", success},
		{"Jenikuuuuuuuuuuuuuu", "Jenik Poustevnicek", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert(Prefix(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestTrie_InsertDuplicatePrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Pepan", "Pepan Zdepan", failure},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert(Prefix(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestTrie_InsertVariousPrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Pepin", "Pepin Omacka", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
		{"Pepan", "Pepan Dupan", failure},
		{"Karel", "Karel Pekar", success},
		{"Jenik", "Jenik Poustevnicek", failure},
		{"Pepanek", "Pepanek Zemlicka", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert(Prefix(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestTrie_InsertAndForceDenseTrie(t *testing.T) {
	var data = []byte(`
C:\
C:\$Recycle.Bin
C:\$Recycle.Bin\S-1-5-21-4152498002-1650887818-552201016-1000
C:\$Recycle.Bin\S-1-5-21-4152498002-1650887818-552201016-1000\desktop.ini
C:\Documents and Settings
C:\Games
C:\Games\Battle.net
C:\Games\Battle.net\.agent.db
C:\Games\Battle.net\.patch.result
C:\Games\Battle.net\Battle.net Launcher.exe
C:\Games\Battle.net\Battle.net.3968
C:\Games\Battle.net\Battle.net.3968\Battle.net.exe
C:\Games\Battle.net\Battle.net.3968\Battle.net.mpq
C:\Games\Battle.net\Battle.net.3968\QtCore4.dll
C:\Games\Battle.net\Battle.net.3968\QtDeclarative4.dll
C:\Games\Battle.net\Battle.net.3968\QtGui4.dll
C:\Games\Battle.net\Battle.net.3968\QtNetwork4.dll
C:\Games\Battle.net\Battle.net.3968\QtScript4.dll
C:\Games\Battle.net\Battle.net.3968\QtSql4.dll
C:\Games\Battle.net\Battle.net.3968\QtXml4.dll
C:\Games\Battle.net\Battle.net.3968\QtXmlPatterns4.dll
C:\Games\Battle.net\Battle.net.3968\battle.net.dll
C:\Games\Battle.net\Battle.net.3968\d3dcompiler_43.dll
C:\Games\Battle.net\Battle.net.3968\d3dx9_43.dll
C:\Games\Battle.net\Battle.net.3968\ffmpegsumo.dll
C:\Games\Battle.net\Battle.net.3968\icudt.dll
C:\Games\Battle.net\Battle.net.3968\imageformats
C:\Games\Battle.net\Battle.net.3968\imageformats\qgif4.dll
C:\Games\Battle.net\Battle.net.3968\imageformats\qico4.dll
C:\Games\Battle.net\Battle.net.3968\imageformats\qjpeg4.dll
C:\Games\Battle.net\Battle.net.3968\imageformats\qmng4.dll
C:\Games\Battle.net\Battle.net.3968\imageformats\qsvg4.dll
C:\Games\Battle.net\Battle.net.3968\imageformats\qtiff4.dll
C:\Games\Battle.net\Battle.net.3968\libEGL.dll
C:\Games\Battle.net\Battle.net.3968\libGLESv2.dll
C:\Games\Battle.net\Battle.net.3968\libcef.dll
C:\Games\Battle.net\Battle.net.3968\locales
C:\Games\Battle.net\Battle.net.3968\locales\de.pak
C:\Games\Battle.net\Battle.net.3968\locales\en-GB.pak
C:\Games\Battle.net\Battle.net.3968\locales\en-US.pak
C:\Games\Battle.net\Battle.net.3968\locales\es.pak
C:\Games\Battle.net\Battle.net.3968\locales\fr.pak
C:\Games\Battle.net\Battle.net.3968\locales\it.pak
C:\Games\Battle.net\Battle.net.3968\locales\ko.pak
C:\Games\Battle.net\Battle.net.3968\locales\pt-BR.pak
C:\Games\Battle.net\Battle.net.3968\locales\pt-PT.pak
C:\Games\Battle.net\Battle.net.3968\locales\ru.pak
C:\Games\Battle.net\Battle.net.3968\locales\zh-CN.pak
C:\Games\Battle.net\Battle.net.3968\locales\zh-TW.pak
C:\Games\Battle.net\Battle.net.3968\msvcp100.dll
C:\Games\Battle.net\Battle.net.3968\msvcr100.dll
C:\Games\Battle.net\Battle.net.4047
C:\Games\Battle.net\Battle.net.4047\Battle.net.exe
C:\Games\Battle.net\Battle.net.4047\Battle.net.mpq
C:\Games\Battle.net\Battle.net.4047\QtCore4.dll
C:\Games\Battle.net\Battle.net.4047\QtDeclarative4.dll
C:\Games\Battle.net\Battle.net.4047\QtGui4.dll
C:\Games\Battle.net\Battle.net.4047\QtNetwork4.dll
C:\Games\Battle.net\Battle.net.4047\QtScript4.dll
C:\Games\Battle.net\Battle.net.4047\QtSql4.dll
C:\Games\Battle.net\Battle.net.4047\QtXml4.dll
C:\Games\Battle.net\Battle.net.4047\QtXmlPatterns4.dll
C:\Games\Battle.net\Battle.net.4047\battle.net.dll
C:\Games\Battle.net\Battle.net.4047\d3dcompiler_43.dll
C:\Games\Battle.net\Battle.net.4047\d3dx9_43.dll
C:\Games\Battle.net\Battle.net.4047\ffmpegsumo.dll
C:\Games\Battle.net\Battle.net.4047\icudt.dll
C:\Games\Battle.net\Battle.net.4047\imageformats
C:\Games\Battle.net\Battle.net.4047\imageformats\qgif4.dll
C:\Games\Battle.net\Battle.net.4047\imageformats\qico4.dll
C:\Games\Battle.net\Battle.net.4047\imageformats\qjpeg4.dll
C:\Games\Battle.net\Battle.net.4047\imageformats\qmng4.dll
C:\Games\Battle.net\Battle.net.4047\imageformats\qsvg4.dll
C:\Games\Battle.net\Battle.net.4047\imageformats\qtiff4.dll
C:\Games\Battle.net\Battle.net.4047\libEGL.dll
C:\Games\Battle.net\Battle.net.4047\libGLESv2.dll
C:\Games\Battle.net\Battle.net.4047\libcef.dll
C:\Games\Battle.net\Battle.net.4047\locales
C:\Games\Battle.net\Battle.net.4047\locales\de.pak
C:\Games\Battle.net\Battle.net.4047\locales\en-GB.pak
C:\Games\Battle.net\Battle.net.4047\locales\en-US.pak
C:\Games\Battle.net\Battle.net.4047\locales\es.pak
C:\Games\Battle.net\Battle.net.4047\locales\fr.pak
C:\Games\Battle.net\Battle.net.4047\locales\it.pak
C:\Games\Battle.net\Battle.net.4047\locales\ko.pak
C:\Games\Battle.net\Battle.net.4047\locales\pt-BR.pak
C:\Games\Battle.net\Battle.net.4047\locales\pt-PT.pak
C:\Games\Battle.net\Battle.net.4047\locales\ru.pak
C:\Games\Battle.net\Battle.net.4047\locales\zh-CN.pak
C:\Games\Battle.net\Battle.net.4047\locales\zh-TW.pak
C:\Games\Battle.net\Battle.net.4047\msvcp100.dll
C:\Games\Battle.net\Battle.net.4047\msvcr100.dll
C:\Games\Battle.net\Battle.net.exe
C:\Games\Battle.net\BlizzardError.exe
C:\Games\Battle.net\Launcher.db
C:\Games\Battle.net\Logs
C:\Games\Battle.net\Logs\Battle.net Install Log.html
C:\Games\Battle.net\Logs\Blizzard Updater Log.html
C:\Games\Battle.net\Logs\Switcher.log
C:\Games\Battle.net\SetupWin.mpq
C:\Games\Battle.net\SystemSurvey.exe
C:\Games\Diablo III
C:\Games\Diablo III\.agent.db
C:\Games\Diablo III\.patch.result
C:\Games\Diablo III\BattlenetAccount.url
C:\Games\Diablo III\Bnet
C:\Games\Diablo III\Bnet\battle.net.dll
C:\Games\Diablo III\D3Debug.txt
C:\Games\Diablo III\Data_D3
C:\Games\Diablo III\Data_D3\PC
C:\Games\Diablo III\Data_D3\PC\MPQs
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\Win
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\Win\patch-Win-10057.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\Win\patch-Win-11327.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\Win\patch-Win-12480.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\Win\patch-Win-16416.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\Win\patch-Win-9558.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-10057.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-10235.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-10485.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-11327.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-12480.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-12811.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-13300.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-13644.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-14633.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-15295.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-16416.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-16603.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-9558.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-9749.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-9858.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-9950.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\base\patch-base-9991.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-10057.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-10485.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-11327.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-12480.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-13300.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-14633.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-16416.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-9558.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-9749.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-9858.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-9950.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Cache\enUS\patch-enUS-9991.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\ClientData.mpq
C:\Games\Diablo III\Data_D3\PC\MPQs\CoreData.mpq
C:\Games\Diablo III\Data_D3\PC\MPQs\HLSLShaders.mpq
C:\Games\Diablo III\Data_D3\PC\MPQs\Sound.mpq
C:\Games\Diablo III\Data_D3\PC\MPQs\Texture.mpq
C:\Games\Diablo III\Data_D3\PC\MPQs\Win
C:\Games\Diablo III\Data_D3\PC\MPQs\Win\d3-update-Win-10057.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Win\d3-update-Win-11327.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Win\d3-update-Win-12480.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Win\d3-update-Win-16416.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\Win\d3-update-Win-9558.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-10057.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-10235.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-10485.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-11327.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-12480.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-12811.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-13300.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-13644.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-14633.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-15295.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-16416.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-16603.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-9558.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-9749.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-9858.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-9950.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base\d3-update-base-9991.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\base-Win.mpq
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-10057.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-10485.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-11327.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-12480.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-13300.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-14633.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-16416.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-9558.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-9749.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-9858.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-9950.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS\d3-update-enUS-9991.MPQ
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS_Audio.mpq
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS_Cutscene.mpq
C:\Games\Diablo III\Data_D3\PC\MPQs\enUS_Text.mpq
C:\Games\Diablo III\Data_D3\PC\realmlist.dtf
C:\Games\Diablo III\Diablo III Launcher.exe
C:\Games\Diablo III\Diablo III.exe
C:\Games\Diablo III\Diablo III.mfil
C:\Games\Diablo III\Diablo III.tfil
C:\Games\Diablo III\InspectorReporter
C:\Games\Diablo III\InspectorReporter\BlizzardError.exe
C:\Games\Diablo III\Launcher.db
C:\Games\Diablo III\Logs
C:\Games\Diablo III\Logs\Blizzard Updater Log.html
C:\Games\Diablo III\Logs\Diablo III Install Log.html
C:\Games\Diablo III\Manual.url
C:\Games\Diablo III\Screenshot19295.png
C:\Games\Diablo III\Screenshot426.png
C:\Games\Diablo III\Screenshot952380.png
C:\Games\Diablo III\Screenshot952381.png
C:\Games\Diablo III\SetupWin.mpq
C:\Games\Diablo III\TechSupport.url
C:\Games\Diablo III\Updates`)

	trie := NewTrie()
	value := true
	for _, line := range bytes.Split(data, []byte("\n")) {
		t.Logf("INSERT prefix=%q, item=%v", line, value)
		if !trie.Insert(Prefix(line), value) {
			t.Errorf("Failed to insert %q", line)
		}
	}
}

func TestTrie_SetGet(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Pepin", "Pepin Omacka", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
		{"Pepan", "Pepan Dupan", failure},
		{"Karel", "Karel Pekar", success},
		{"Jenik", "Jenik Poustevnicek", failure},
		{"Pepanek", "Pepanek Zemlicka", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert(Prefix(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range data {
		t.Logf("SET %q to 10", v.key)
		trie.Set(Prefix(v.key), 10)
	}

	for _, v := range data {
		value := trie.Get(Prefix(v.key))
		t.Logf("GET %q => %v", v.key, value)
		if value.(int) != 10 {
			t.Errorf("Unexpected return value, != 10", value)
		}
	}

	if value := trie.Get(Prefix("random crap")); value != nil {
		t.Errorf("Unexpected return value, %v != <nil>", value)
	}
}

func TestTrie_Match(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Pepin", "Pepin Omacka", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
		{"Pepan", "Pepan Dupan", failure},
		{"Karel", "Karel Pekar", success},
		{"Jenik", "Jenik Poustevnicek", failure},
		{"Pepanek", "Pepanek Zemlicka", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert(Prefix(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range data {
		matched := trie.Match(Prefix(v.key))
		t.Logf("MATCH %q => %v", v.key, matched)
		if !matched {
			t.Errorf("Inserted key %q was not matched", v.key)
		}
	}

	if trie.Match(Prefix("random crap")) {
		t.Errorf("Key that was not inserted matched: %q", "random crap")
	}
}

func TestTrie_MatchSubtree(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Pepin", "Pepin Omacka", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
		{"Pepan", "Pepan Dupan", failure},
		{"Karel", "Karel Pekar", success},
		{"Jenik", "Jenik Poustevnicek", failure},
		{"Pepanek", "Pepanek Zemlicka", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert(Prefix(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range data {
		key := Prefix(v.key[:3])
		matched := trie.MatchSubtree(key)
		t.Logf("MATCH_SUBTREE %q => %v", key, matched)
		if !matched {
			t.Errorf("Subtree %q was not matched", v.key)
		}
	}
}

func TestTrie_Visit(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepa", 0, success},
		{"Pepa Zdepa", 1, success},
		{"Pepa Kuchar", 2, success},
		{"Honza", 3, success},
		{"Jenik", 4, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	if err := trie.Visit(func(prefix Prefix, item Item) error {
		name := data[item.(int)].key
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if !strings.HasPrefix(string(prefix), name) {
			t.Errorf("Unexpected prefix encountered, %q not a prefix of %q", prefix, name)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestTrie_VisitSkipSubtree(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepa", 0, success},
		{"Pepa Zdepa", 1, success},
		{"Pepa Kuchar", 2, success},
		{"Honza", 3, success},
		{"Jenik", 4, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	if err := trie.Visit(func(prefix Prefix, item Item) error {
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if item.(int) == 0 {
			t.Logf("SKIP %q", prefix)
			return SkipSubtree
		}
		if strings.HasPrefix(string(prefix), "Pepa") {
			t.Errorf("Unexpected prefix encountered, %q", prefix)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestTrie_VisitReturnError(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepa", 0, success},
		{"Pepa Zdepa", 1, success},
		{"Pepa Kuchar", 2, success},
		{"Honza", 3, success},
		{"Jenik", 4, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	someErr := errors.New("Something exploded")
	if err := trie.Visit(func(prefix Prefix, item Item) error {
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if item.(int) == 0 {
			return someErr
		}
		if item.(int) != 0 {
			t.Errorf("Unexpected prefix encountered, %q", prefix)
		}
		return nil
	}); err != nil && err != someErr {
		t.Fatal(err)
	}
}

func TestTrie_VisitSubtree(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepa", 0, success},
		{"Pepa Zdepa", 1, success},
		{"Pepa Kuchar", 2, success},
		{"Honza", 3, success},
		{"Jenik", 4, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	var counter int
	subtreePrefix := []byte("Pep")
	t.Log("VISIT Pep")
	if err := trie.VisitSubtree(subtreePrefix, func(prefix Prefix, item Item) error {
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if !bytes.HasPrefix(prefix, subtreePrefix) {
			t.Errorf("Unexpected prefix encountered, %q does not extend %q",
				prefix, subtreePrefix)
		}
		if len(prefix) > len(data[item.(int)].key) {
			t.Fatalf("Something is rather fishy here, prefix=%q", prefix)
		}
		counter++
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if counter != 3 {
		t.Error("Unexpected number of nodes visited")
	}
}

func TestTrie_VisitPrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"P", 0, success},
		{"Pe", 1, success},
		{"Pep", 2, success},
		{"Pepa", 3, success},
		{"Pepa Zdepa", 4, success},
		{"Pepa Kuchar", 5, success},
		{"Honza", 6, success},
		{"Jenik", 7, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	var counter int
	word := []byte("Pepa")
	if err := trie.VisitPrefixes(word, func(prefix Prefix, item Item) error {
		t.Logf("VISITING prefix=%q, item=%v", prefix, item)
		if !bytes.HasPrefix(word, prefix) {
			t.Errorf("Unexpected prefix encountered, %q is not a prefix of %q",
				prefix, word)
		}
		counter++
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if counter != 4 {
		t.Error("Unexpected number of nodes visited")
	}
}

func TestParticiaTrie_Delete(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range data {
		t.Logf("DELETE word=%v, success=%v", v.key, v.retVal)
		if ok := trie.Delete([]byte(v.key)); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestParticiaTrie_DeleteNonExistent(t *testing.T) {
	trie := NewTrie()

	insertData := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Honza", "Honza Novak", success},
		{"Jenik", "Jenik Poustevnicek", success},
	}
	deleteData := []testData{
		{"Pepan", "Pepan Zdepan", success},
		{"Honza", "Honza Novak", success},
		{"Pepan", "Pepan Zdepan", failure},
		{"Jenik", "Jenik Poustevnicek", success},
		{"Honza", "Honza Novak", failure},
	}

	for _, v := range insertData {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range deleteData {
		t.Logf("DELETE word=%v, success=%v", v.key, v.retVal)
		if ok := trie.Delete([]byte(v.key)); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestParticiaTrie_DeleteSubtree(t *testing.T) {
	trie := NewTrie()

	insertData := []testData{
		{"P", 0, success},
		{"Pe", 1, success},
		{"Pep", 2, success},
		{"Pepa", 3, success},
		{"Pepa Zdepa", 4, success},
		{"Pepa Kuchar", 5, success},
		{"Honza", 6, success},
		{"Jenik", 7, success},
	}
	deleteData := []testData{
		{"Pe", -1, success},
		{"Pe", -1, failure},
		{"Honzik", -1, failure},
		{"Honza", -1, success},
		{"Honza", -1, failure},
		{"Pep", -1, failure},
		{"P", -1, success},
		{"Nobody", -1, failure},
		{"", -1, success},
	}

	for _, v := range insertData {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Fatalf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}

	for _, v := range deleteData {
		t.Logf("DELETE_SUBTREE prefix=%v, success=%v", v.key, v.retVal)
		if ok := trie.DeleteSubtree([]byte(v.key)); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

/*
func TestTrie_Dump(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"Honda", nil, success},
		{"Honza", nil, success},
		{"Jenik", nil, success},
		{"Pepan", nil, success},
		{"Pepin", nil, success},
	}

	for i, v := range data {
		if _, ok := trie.Insert([]byte(v.key), v.value); ok != v.retVal {
			t.Logf("INSERT %v %v", v.key, v.value)
			t.Fatalf("Unexpected return value, expected=%v, got=%v", i, ok)
		}
	}

	dump := `
+--+--+ Hon +--+--+ da
   |           |
   |           +--+ za
   |
   +--+ Jenik
   |
   +--+ Pep +--+--+ an
               |
               +--+ in
`

	var buf bytes.Buffer
	trie.Dump(buf)

	if !bytes.Equal(buf.Bytes(), dump) {
		t.Logf("DUMP")
		t.Fatalf("Unexpected dump generated, expected\n\n%v\ngot\n\n%v", dump, buf.String())
	}
}
*/

func TestTrie_compact(t *testing.T) {
	trie := NewTrie()

	trie.Insert(Prefix("a"), 0)
	trie.Insert(Prefix("ab"), 0)
	trie.Insert(Prefix("abc"), 0)
	trie.Insert(Prefix("abcd"), 0)
	trie.Insert(Prefix("abcde"), 0)
	trie.Insert(Prefix("abcdef"), 0)
	trie.Insert(Prefix("abcdefg"), 0)
	trie.Insert(Prefix("abcdefgi"), 0)
	trie.Insert(Prefix("abcdefgij"), 0)
	trie.Insert(Prefix("abcdefgijk"), 0)

	trie.Delete(Prefix("abcdef"))
	trie.Delete(Prefix("abcde"))
	trie.Delete(Prefix("abcdefg"))

	trie.Delete(Prefix("a"))
	trie.Delete(Prefix("abc"))
	trie.Delete(Prefix("ab"))

	trie.Visit(func(prefix Prefix, item Item) error {
		// 97 ~~ 'a',
		for ch := byte(97); ch <= 107; ch++ {
			if c := bytes.Count(prefix, []byte{ch}); c > 1 {
				t.Errorf("%q appeared in %q %v times", ch, prefix, c)
			}
		}
		return nil
	})
}

func TestTrie_longestCommonPrefixLenght(t *testing.T) {
	trie := NewTrie()
	trie.prefix = []byte("1234567890")

	switch {
	case trie.longestCommonPrefixLength([]byte("")) != 0:
		t.Fail()
	case trie.longestCommonPrefixLength([]byte("12345")) != 5:
		t.Fail()
	case trie.longestCommonPrefixLength([]byte("123789")) != 3:
		t.Fail()
	case trie.longestCommonPrefixLength([]byte("12345678901")) != 10:
		t.Fail()
	}
}

// Examples --------------------------------------------------------------------

func ExampleTrie() {
	// Create a new tree.
	trie := NewTrie()

	// Insert some items.
	trie.Insert(Prefix("Pepa Novak"), 1)
	trie.Insert(Prefix("Pepa Sindelar"), 2)
	trie.Insert(Prefix("Karel Macha"), 3)
	trie.Insert(Prefix("Karel Hynek Macha"), 4)

	// Just check if some things are present in the tree.
	key := Prefix("Pepa Novak")
	fmt.Printf("%q present? %v\n", key, trie.Match(key))
	key = Prefix("Karel")
	fmt.Printf("Anybody called %q here? %v\n", key, trie.MatchSubtree(key))

	// Walk the tree.
	trie.Visit(printItem)
	// "Pepa Novak": 1
	// "Pepa Sindelar": 2
	// "Karel Macha": 3
	// "Karel Hynek Macha": 4

	// Walk a subtree.
	trie.VisitSubtree(Prefix("Pepa"), printItem)
	// "Pepa Novak": 1
	// "Pepa Sindelar": 2

	// Modify an item, then fetch it from the tree.
	trie.Set(Prefix("Karel Hynek Macha"), 10)
	key = Prefix("Karel Hynek Macha")
	fmt.Printf("%q: %v\n", key, trie.Get(key))
	// "Karel Hynek Macha": 10

	// Walk prefixes.
	prefix := Prefix("Karel Hynek Macha je kouzelnik")
	trie.VisitPrefixes(prefix, printItem)
	// "Karel Hynek Macha": 10

	// Delete some items.
	trie.Delete(Prefix("Pepa Novak"))
	trie.Delete(Prefix("Karel Macha"))

	// Walk again.
	trie.Visit(printItem)
	// "Pepa Sindelar": 2
	// "Karel Hynek Macha": 10

	// Delete a subtree.
	trie.DeleteSubtree(Prefix("Pepa"))

	// Print what is left.
	trie.Visit(printItem)
	// "Karel Hynek Macha": 10

	// Output:
	// "Pepa Novak" present? true
	// Anybody called "Karel" here? true
	// "Pepa Novak": 1
	// "Pepa Sindelar": 2
	// "Karel Macha": 3
	// "Karel Hynek Macha": 4
	// "Pepa Novak": 1
	// "Pepa Sindelar": 2
	// "Karel Hynek Macha": 10
	// "Karel Hynek Macha": 10
	// "Pepa Sindelar": 2
	// "Karel Hynek Macha": 10
	// "Karel Hynek Macha": 10
}

// Helpers ---------------------------------------------------------------------

func printItem(prefix Prefix, item Item) error {
	fmt.Printf("%q: %v\n", prefix, item)
	return nil
}
