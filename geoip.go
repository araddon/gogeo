package gogeo
/*
#cgo LDFLAGS: -lGeoIP
#include <stdio.h>
#include <errno.h>
#include <GeoIPCity.h>
*/
import "C"
import "unsafe"
import "errors"
import "net"

type GeoIPFlag int
var (
    Standard = 0
    MemoryCache = 1
    CheckCache = 2
    IndexCache = 2
    MMapCache = 4
    )

type GeoIP struct {
    GeoIP *C.GeoIP
}

type GeoIPRecord struct {
    CountryCode string
    CountryCode3 string
    CountryName string
    Region string
    City string
    PostalCode string
    Latitude float64
    Longitude float64
    AreaCode int
    CharSet int
    ContinentCode string
    CountryConfidence byte
    RegionConfidence byte
    CityConfidence byte
    PostalConfidence byte
    AccuracyRadius int
}
func parseGeoIPRecord(c_record *C.GeoIPRecord) *GeoIPRecord {
    if c_record == nil {
        return nil
    }
    record := new(GeoIPRecord)
    record.CountryCode = C.GoString(c_record.country_code)
    record.CountryCode3 = C.GoString(c_record.country_code3)
    record.CountryName = C.GoString(c_record.country_name)
    record.Region = C.GoString(c_record.region)
    record.PostalCode = C.GoString(c_record.postal_code)
    record.Latitude = float64(c_record.latitude)
    record.Longitude = float64(c_record.longitude)
    record.AreaCode = int(c_record.area_code)
    record.CharSet = int(c_record.charset)
    record.ContinentCode = C.GoString(c_record.continent_code)
    record.CountryConfidence = byte(c_record.country_conf)
    record.RegionConfidence = byte(c_record.region_conf)
    record.CityConfidence = byte(c_record.city_conf)
    record.PostalConfidence = byte(c_record.postal_conf)
    record.AccuracyRadius = int(c_record.accuracy_radius)
    return record
}
func Open(filename string, flags int) (*GeoIP, error) {
    base := C.CString(filename)
    defer C.free(unsafe.Pointer(base))
    db := C.GeoIP_open(base, C.int(flags))
    if db == nil {
        return nil, errors.New("Cannot create GeoIP object")
    }
    geoIP := &GeoIP{db}
    return geoIP, nil
}

func (self *GeoIP) Close() {
    if self.GeoIP != nil {
        C.GeoIP_delete(self.GeoIP)
        self.GeoIP = nil
    }
}

func IPv4ToInt(ip []byte) uint32 {
    var ipaddr uint32
    ipaddr |= uint32(ip[0]) << 24
    ipaddr |= uint32(ip[1]) << 16
    ipaddr |= uint32(ip[2]) << 8
    ipaddr |= uint32(ip[3])
    return ipaddr
}

func (self *GeoIP) RecordByIPAddr(addr *net.IPAddr) *GeoIPRecord {
    ip := addr.IP
    if len(ip) == 4 {
        ipaddr := IPv4ToInt(ip)
        record := C.GeoIP_record_by_ipnum(self.GeoIP, C.ulong(ipaddr))
        defer C.GeoIPRecord_delete(record)
        return parseGeoIPRecord(record)
    } else if len(ip) == 16 {
    }
    return nil
}
