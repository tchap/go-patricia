// Copyright (c) 2014 The go-patricia AUTHORS
//
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package patricia

import (
	"bytes"
	"testing"
)

// Tests -----------------------------------------------------------------------

func TestTrie_InsertDense(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"aba", 0, success},
		{"abb", 1, success},
		{"abc", 2, success},
		{"abd", 3, success},
		{"abe", 4, success},
		{"abf", 5, success},
		{"abg", 6, success},
		{"abh", 7, success},
		{"abi", 8, success},
		{"abj", 9, success},
		{"abk", 0, success},
		{"abl", 1, success},
		{"abm", 2, success},
		{"abn", 3, success},
		{"abo", 4, success},
		{"abp", 5, success},
		{"abq", 6, success},
		{"abr", 7, success},
		{"abs", 8, success},
		{"abt", 9, success},
		{"abu", 0, success},
		{"abv", 1, success},
		{"abw", 2, success},
		{"abx", 3, success},
		{"aby", 4, success},
		{"abz", 5, success},
	}

	for _, v := range data {
		t.Logf("INSERT prefix=%v, item=%v, success=%v", v.key, v.value, v.retVal)
		if ok := trie.Insert(Prefix(v.key), v.value); ok != v.retVal {
			t.Errorf("Unexpected return value, expected=%v, got=%v", v.retVal, ok)
		}
	}
}

func TestTrie_InsertDenseDuplicatePrefixes(t *testing.T) {
	trie := NewTrie()

	data := []testData{
		{"aba", 0, success},
		{"abb", 1, success},
		{"abc", 2, success},
		{"abd", 3, success},
		{"abe", 4, success},
		{"abf", 5, success},
		{"abg", 6, success},
		{"abh", 7, success},
		{"abi", 8, success},
		{"abj", 9, success},
		{"abk", 0, success},
		{"abl", 1, success},
		{"abm", 2, success},
		{"abn", 3, success},
		{"abo", 4, success},
		{"abp", 5, success},
		{"abq", 6, success},
		{"abr", 7, success},
		{"abs", 8, success},
		{"abt", 9, success},
		{"abu", 0, success},
		{"abv", 1, success},
		{"abw", 2, success},
		{"abx", 3, success},
		{"aby", 4, success},
		{"abz", 5, success},
		{"aba", 0, failure},
		{"abb", 1, failure},
		{"abc", 2, failure},
		{"abd", 3, failure},
		{"abe", 4, failure},
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
