package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.2" )

func 0N5iGs0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q226GGGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrVXr1BmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8Y0jt0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lX8IYZzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbEegIQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oKsduubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5PWq9GZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wI9w4o6jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8j3OYteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvlwknZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOShiq8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEDHmpQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJDEMfsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJnXxGVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2a0UHZKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrJmIoCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJHROmCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CFO22ZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eDTZ0RvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEIhpCpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWFdNN2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faZbW9LRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBZ1OrWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJRDgWKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPVn7x92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vhx6fzvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lYd4h9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnGfc2NGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3x3ARsHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJLO1JFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8grY5RLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCL3C0VZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWkpgNz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBn9EHTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1TNkFvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mprt6dYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dupQIH98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8r7KGRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVaizwMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTTUrqvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZblJAEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxYGNyAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k13Ucu3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFoBuRr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjMFIXcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKkRs4YQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzY5NMSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOG7Iq7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lF4HXSIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLZh1exGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxQoNSJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHI8xjmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5C4wb5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7Z9Pj8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBE0CFL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LMK31dJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87T5slxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MrfINKH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pK9oiemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaxzAJy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gcZYAvD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUOSiEoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4AJUvIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5utYul0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLXgPikdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lEtmjimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGtCDN8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZJgYdz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUZGfeyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74nME06vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFXDMw9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nCML4eRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YuMOSyUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7h4CfJqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQiSi7rrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4VGZ5DfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BIEP3EWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5VIfx2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1b5iDBsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwK4vHbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBKwZul3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbXkFgZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hhDplBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOyC3uhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZJ0sKq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func By1RCbKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xO5lzSutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOx1BnroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9tDW1A3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wjQv3ItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4k5Rj4rvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsxwV2b4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzH3H6SZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdFJzhLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mOqoFEOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YiMxvZxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trn2qfLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ziDddHxKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rMcxz5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyM17lLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MnHrH6ctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHhQZ0R9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1yhi0FoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdTsEo7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jDZggg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjsCmhBMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3el4W5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Daf6vAEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQWKwhruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcQWzMfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFPedxKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgdyDWmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQnnoIHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIWDwgCfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xaPi15pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDaNvU7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7QDTkE8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQuxumRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsAqlpR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKQccoN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQk5hPO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGOCRafvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ooCHS6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oa43UPNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHz8zRwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXSnVnZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwJMLSGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DWsn5PSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fJnPAfM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gyB12lNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlrr9wkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtZ5fkLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCKHUs33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wb4ygc6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uqVl2SsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCwPDt3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9GS7Q4VnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YROyt88YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E20CBnmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LmDaMZhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASRlRrQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jdYaKa6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYFD4xIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkRPmSEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErugoIzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClyZQV8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4zdOHM6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkgzHOLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LG25xoxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZGVNzYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zF1BohTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfimZFieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iu2a0IVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NN4vevlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgKwbElqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbldGvT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgPNKtBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0yHbALaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiGNMkTVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eixAtfCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwFyVQzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEbcWzvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auDrje6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFC0sYjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHfL9EmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wr8B6bTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhBF3gXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfcf7vkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLtdSMeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gILq9hUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z91SdVYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBoybG8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TssycohVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5o94awQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bw8WR4B0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z49NgWl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTAyAUcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vn6xZ88BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwXVWTTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRSXvvyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJYjas4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Mj0U3VIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dvUmI4oEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGcjAWfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9g2l9xfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpTwNhIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lSshcAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMrBDrzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORt1aQBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wevBDHfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y19RgEkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBBoXwMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0FdTvv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzXUCr42Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7P5KrKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fC7WVtXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ORh4IzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXC4cl7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 651y7GUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDrPlUUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICbJGGjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOGZ9xGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfdcJoi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzG0JAQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIJldHPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mMXfUALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwJeQ8qxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OA6jIThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoH4WWpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TPch4bjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXV4LlolWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTXFTfebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBcjYGWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5z8uCYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uF8gJeVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xn5mLE4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlKo86NkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHezJi9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cXpkJBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONrKDpcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDNO1e0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwmxiBf9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7dqMcQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOAqVao4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J52VDSydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYmaY1CGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZkeG1pAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c425lgA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fiIByRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpYdl5lCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mm8M68J6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiRPINzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pmu6O0d8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRrCkib5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cjxkZ50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yi23bXE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgtMYMcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5ralAfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rly9xVlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func df4iC2oQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbvWKYUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fEkOhxeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPUUFcBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Zjkbd9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvYwNN6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ll8JF9upWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYnKFkloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSBBdyZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cON8tceIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u998yWCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OItLeBVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZhqvhdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vz8xFv3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLGWcS02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbWQjkZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2VGWcPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZABKXlbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5QNRnfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJEC9pSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcEMQRpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pM6P4cp8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Pw70JjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGbAluxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYIBNEJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16uD5okaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ho19fEIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5UItiKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1Judd9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ue1kD5zjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkplwRC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3t5WkejHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbiTlBDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S01c8RphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rlm4WDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiGlRvDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFdOHMPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaCSW9vPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jkd2Le1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9bqEzb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFf9zTLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYLGbT4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNjYSbG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4yNGIS4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixVhExsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZQ4GJENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ca6jXmwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3j5byu8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRIYBFwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TybphHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e81fOEEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRbdVsuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4ka1SD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2nI30TkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAOp9pGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ugn860aVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcfFCCi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnURDDclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YICwGrlXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEZvq3nhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URD1ERRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiYIEYFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 749EGg2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXl8czGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFW0C9VtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arJl5EjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bC7IYE0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEYIqlfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s91o3RLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qONYSQ9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmMv9E0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnoUAaQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func THtZmJdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 622xZeEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jjh7XNdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOOBOdEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFbXuDsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9uyxeZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6m9TCOnCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8im8LcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yalO8yvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

