/*
 Copyright (c) 2012, Shane Hansen
All rights reserved.

 Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package gogeo

/*
#cgo LDFLAGS: -lGeoIP
#include <stdio.h>
#include <errno.h>
#include <GeoIPCity.h>
*/
import "C"

type GeoIPOptions int

const (
    Standard    = GeoIPOptions(C.GEOIP_STANDARD)
    MemoryCache = GeoIPOptions(C.GEOIP_MEMORY_CACHE)
    CheckCache  = GeoIPOptions(C.GEOIP_CHECK_CACHE)
    IndexCache  = GeoIPOptions(C.GEOIP_INDEX_CACHE)
    MMapCache   = GeoIPOptions(C.GEOIP_MMAP_CACHE)
)

type GeoIPDBType int

const (
    COUNTRY_EDITION                    = GeoIPDBType(C.GEOIP_COUNTRY_EDITION)
    REGION_EDITION_REV0                = GeoIPDBType(C.GEOIP_REGION_EDITION_REV0)
    CITY_EDITION_REV0                  = GeoIPDBType(C.GEOIP_CITY_EDITION_REV0)
    ORG_EDITION                        = GeoIPDBType(C.GEOIP_ORG_EDITION)
    ISP_EDITION                        = GeoIPDBType(C.GEOIP_ISP_EDITION)
    CITY_EDITION_REV1                  = GeoIPDBType(C.GEOIP_CITY_EDITION_REV1)
    REGION_EDITION_REV1                = GeoIPDBType(C.GEOIP_REGION_EDITION_REV1)
    PROXY_EDITION                      = GeoIPDBType(C.GEOIP_PROXY_EDITION)
    ASNUM_EDITION                      = GeoIPDBType(C.GEOIP_ASNUM_EDITION)
    NETSPEED_EDITION                   = GeoIPDBType(C.GEOIP_NETSPEED_EDITION)
    DOMAIN_EDITION                     = GeoIPDBType(C.GEOIP_DOMAIN_EDITION)
    COUNTRY_EDITION_V6                 = GeoIPDBType(C.GEOIP_COUNTRY_EDITION_V6)
    LOCATIONA_EDITION                  = GeoIPDBType(C.GEOIP_LOCATIONA_EDITION)
    ACCURACYRADIUS_EDITION             = GeoIPDBType(C.GEOIP_ACCURACYRADIUS_EDITION)
    CITYCONFIDENCE_EDITION             = GeoIPDBType(C.GEOIP_CITYCONFIDENCE_EDITION)
    CITYCONFIDENCEDIST_EDITION         = GeoIPDBType(C.GEOIP_CITYCONFIDENCEDIST_EDITION)
    LARGE_COUNTRY_EDITION              = GeoIPDBType(C.GEOIP_LARGE_COUNTRY_EDITION)
    LARGE_COUNTRY_EDITION_V6           = GeoIPDBType(C.GEOIP_LARGE_COUNTRY_EDITION_V6)
    CITYCONFIDENCEDIST_ISP_ORG_EDITION = GeoIPDBType(C.GEOIP_CITYCONFIDENCEDIST_ISP_ORG_EDITION)
    CCM_COUNTRY_EDITION                = GeoIPDBType(C.GEOIP_CCM_COUNTRY_EDITION)
    ASNUM_EDITION_V6                   = GeoIPDBType(C.GEOIP_ASNUM_EDITION_V6)
    ISP_EDITION_V6                     = GeoIPDBType(C.GEOIP_ISP_EDITION_V6)
    ORG_EDITION_V6                     = GeoIPDBType(C.GEOIP_ORG_EDITION_V6)
    DOMAIN_EDITION_V6                  = GeoIPDBType(C.GEOIP_DOMAIN_EDITION_V6)
    LOCATIONA_EDITION_V6               = GeoIPDBType(C.GEOIP_LOCATIONA_EDITION_V6)
    REGISTRAR_EDITION                  = GeoIPDBType(C.GEOIP_REGISTRAR_EDITION)
    REGISTRAR_EDITION_V6               = GeoIPDBType(C.GEOIP_REGISTRAR_EDITION_V6)
    USERTYPE_EDITION                   = GeoIPDBType(C.GEOIP_USERTYPE_EDITION)
    USERTYPE_EDITION_V6                = GeoIPDBType(C.GEOIP_USERTYPE_EDITION_V6)
    CITY_EDITION_REV1_V6               = GeoIPDBType(C.GEOIP_CITY_EDITION_REV1_V6)
    CITY_EDITION_REV0_V6               = GeoIPDBType(C.GEOIP_CITY_EDITION_REV0_V6)
    NETSPEED_EDITION_REV1              = GeoIPDBType(C.GEOIP_NETSPEED_EDITION_REV1)
    NETSPEED_EDITION_REV1_V6           = GeoIPDBType(C.GEOIP_NETSPEED_EDITION_REV1_V6)
)
