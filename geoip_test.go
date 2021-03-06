/*
 Copyright (c) 2012, Shane Hansen
All rights reserved.

 Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package gogeo

import "net"
import "testing"

func TestIPv4(t *testing.T) {
    db, err := Open("/usr/share/GeoIP/GeoLiteCity.dat", MemoryCache)
    if err != nil {
        t.Fatal("Coulnd't open database, do you have it installed?")
    }
    if db.DatabaseEdition() != CITY_EDITION_REV1 {
        t.Fatal(`gegeo hasn't been tested with this kind of db,
 proceed at your own risk`)
    }
    addr, err := net.ResolveIPAddr("ip4", "google.com")
    record := db.RecordByIPAddr(addr)
    if record.CountryCode != "US" {
        t.Fatal("it didn't work")
    }
}

func TestIPv6(t *testing.T) {
    db, err := Open("/usr/share/GeoIP/GeoLiteCityIPv6.dat", MemoryCache)
    if err != nil {
        t.Fatal("Coulnd't open database, do you have it installed?")
    }
    if db.DatabaseEdition() != CITY_EDITION_REV1_V6 {
        t.Fatal(`gegeo hasn't been tested with this kind of db,
 proceed at your own risk`)
    }
    addr, err := net.ResolveIPAddr("ip6", "google.com")
    if err != nil {
        t.Fatal(err)
    }
    record := db.RecordByIPAddr(addr)
    if record.CountryCode != "US" {
        t.Fatal("it didn't work")
    }
}
