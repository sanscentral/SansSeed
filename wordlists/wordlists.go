/*
	SansSeed is a BIP39 compatible implementation for generating mnemonic phrases and seed derivation
	Copyright (C) 2018  Sans Central

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package wordlists

import (
	"fmt"
	"strings"

	"github.com/sanscentral/sansseed/wordlists/languages"
)

// BIP39Wordlist is the interface used for BIP39 word lists
// this interface is implemented on the basis of more languages being added in the future
type BIP39Wordlist interface {
	GetName() string
	GetList() []string
	GetChecksum() string
}

// GetAll returns all available word lists
// using this method means that all word lists will be loaded into memory
// it is less memory intensive to instead
// make a call to GetAllWordListNames() followed by a call to GetWordListByName()
func GetAll() []BIP39Wordlist {
	return []BIP39Wordlist{
		languages.BIP39ChineaseSimplified{},
		languages.BIP39ChineaseTraditional{},
		languages.BIP39English{},
		languages.BIP39French{},
		languages.BIP39Italian{},
		languages.BIP39Japanese{},
		languages.BIP39Korean{},
		languages.BIP39Spanish{},
	}
}

// GetAllLanguageNames returns list of all available languages
func GetAllLanguageNames() []string {
	r := []string{}
	for _, l := range GetAll() {
		r = append(r, l.GetName())
	}
	return r
}

// GetByLanguageName returns a wordlist by its language name
// 'n' must match an entry from GetAllLanguageNames()
// comparison is case insensitive
func GetByLanguageName(n string) (*BIP39Wordlist, error) {
	nl := strings.ToLower(n)
	for _, l := range GetAll() {
		if strings.ToLower(l.GetName()) == nl {
			return &l, nil
		}
	}
	return nil, fmt.Errorf("no matching wordlist for language '%s' was found", n)
}
