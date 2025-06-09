package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
	"github.com/samborkent/uuidv7"
	"github.com/segmentio/ksuid"
)

func main() {
	var (
		count   int
		idType  string
		nanoLen int
	)

	flag.IntVar(&count, "n", 1, "Number of IDs to generate")
	flag.StringVar(&idType, "type", "ulid", "ID type: ulid | uuid4 | uuid5 | uuid6 | uuid7 | ksuid | xid | nanoid | snowflake")
	flag.IntVar(&nanoLen, "len", 21, "NanoID length (applies only to nanoid)")
	flag.Parse()

	if count < 1 {
		log.Fatalf("Count must be at least 1")
	}

	switch idType {
	case "ulid":
		generateULIDs(count)
	case "uuid4":
		generateUUID4(count)
	case "uuid5":
		generateUUID5(count)
	case "uuid6":
		generateUUID6(count)
	case "uuid7":
		generateUUID7(count)
	case "ksuid":
		generateKSUIDs(count)
	case "xid":
		generateXIDs(count)
	case "nanoid":
		generateNanoIDs(count, nanoLen)
	case "snowflake":
		generateSnowflakes(count)
	default:
		log.Fatalf("Unsupported ID type: %s", idType)
	}
}

func generateUUID4(n int) {
	for i := 0; i < n; i++ {
		id := uuid.Must(uuid.NewRandom())
		fmt.Println(id.String())
	}
}

func generateUUID5(n int) {
	ns := uuid.NameSpaceDNS
	for i := 0; i < n; i++ {
		id := uuid.NewSHA1(ns, []byte(fmt.Sprintf("example-%d", i)))
		fmt.Println(id.String())
	}
}

func generateUUID6(n int) {
	for i := 0; i < n; i++ {
		t := time.Now().UnixMilli()
		u := uuid.New()
		bytes := u[:]
		bytes[0] = byte(t >> 24)
		bytes[1] = byte(t >> 16)
		bytes[2] = byte(t >> 8)
		bytes[3] = byte(t)
		bytes[6] = (bytes[6] & 0x0f) | 0x60 // version 6
		fmt.Printf("%x-%x-%x-%x-%x\n", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
	}
}

func generateUUID7(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(uuidv7.New().String())
	}
}

func generateULIDs(n int) {
	entropy := ulid.Monotonic(rand.Reader, math.MaxInt64)
	for i := 0; i < n; i++ {
		id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
		fmt.Println(id.String())
	}
}

func generateKSUIDs(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(ksuid.New().String())
	}
}

func generateXIDs(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(xid.New().String())
	}
}

func generateNanoIDs(n int, length int) {
	for i := 0; i < n; i++ {
		id, err := gonanoid.New(length)
		if err != nil {
			log.Fatalf("Failed to generate NanoID: %v", err)
		}
		fmt.Println(id)
	}
}

func generateSnowflakes(n int) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatalf("Failed to initialise snowflake node: %v", err)
	}
	for i := 0; i < n; i++ {
		fmt.Println(node.Generate().String())
	}
}
