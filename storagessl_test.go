package storagessl

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestNats_Exists(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "MultiComDatabase"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	bln := A.Exists(ctx, "test")
	if bln {
		fmt.Println("El Certificado en la base " + A.DataBaseName + ": Existe")
	} else {
		fmt.Println("El Certificado en la base " + A.DataBaseName + ": NO Existe")
	}
}

func TestNats_Store(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "MultiComDatabase"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	aByte, _ := ioutil.ReadFile("C:\\Application\\Amazon PEM\\loadbalancer\\loadbalancer.reg")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	A.Store(ctx, "test", aByte)
}

func TestNats_Store2(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "mcom_adnud_dev"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	aByte, _ := ioutil.ReadFile("C:\\Application\\Amazon PEM\\loadbalancer\\serverpronto.nodo-4.ppk")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	A.Store(ctx, "test", aByte)
}

func TestNats_Store3(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "MultiComDatabase"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	aByte, _ := ioutil.ReadFile("C:\\Application\\Amazon PEM\\loadbalancer\\amazon_load_balancer.pem")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	A.Store(ctx, "rudo", aByte)

	aByte1, _ := ioutil.ReadFile("C:\\Application\\Amazon PEM\\loadbalancer\\amazon_load_balancer.pem")
	A.Store(ctx, "alex", aByte1)
	A.Store(ctx, "victor", aByte1)
	A.Store(ctx, "arizona", aByte1)
	A.Store(ctx, "america", aByte1)
	A.Store(ctx, "anahuac", aByte1)
	A.Store(ctx, "analiza", aByte1)
}

func TestNats_Load(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "MultiComDatabase"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := A.Load(ctx, "test")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Buffer loaded from " + A.DataBaseName)
	}
}

func TestNats_Load2(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "mcom_adnud_dev"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := A.Load(ctx, "test")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Buffer loaded from " + A.DataBaseName)
	}
}

func TestNats_Exists2(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "MultiComDatabase"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	bln := A.Exists(ctx, "test")
	if bln {
		fmt.Println("El Certificado en la base " + A.DataBaseName + ": Existe")
	} else {
		fmt.Println("El Certificado en la base " + A.DataBaseName + ": NO Existe")
	}
}

func TestNats_Delete(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "MultiComDatabase"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err := A.Delete(ctx, "test")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Buffer delete from " + A.DataBaseName)
	}
}

func TestNats_Delete2(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "mcom_adnud_dev"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err := A.Delete(ctx, "test")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Buffer delete from " + A.DataBaseName)
	}
}

func TestNats_Exists3(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "MultiComDatabase"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	bln := A.Exists(ctx, "test")
	if bln {
		fmt.Println("El Certificado en la base " + A.DataBaseName + ": Existe")
	} else {
		fmt.Println("El Certificado en la base " + A.DataBaseName + ": NO Existe")
	}
}

func TestNats_Stat(t *testing.T) {
	A := new(StorageParam)
	A.Host = "127.0.0.1"
	A.Port = "9582"
	A.DataBaseName = "MultiComDatabase"
	A.Timeout = 5
	A.AccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	A.Stat(ctx, "victor")
}
