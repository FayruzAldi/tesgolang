// package main

// import SayHai "golang-basic/SayHai"
//  // Pastikan modul ini ada di go.mod

// func j() {
// 	// var dataFloat float32 = 3.50
// 	// var parseToInt int = int(dataFloat)	
// 	// var x bool = true
// 	// var nilai int = 81
// 	// sum := 0
// 	// var y [5]int

// 	// y[0] = 30
// 	// y[1] = 66
// 	// y[2] = 75
// 	// y[3] = 99
// 	// y[4] = 100

// 	// var array = [...]int{
// 	// 	1,2,3,4,5,6,7,8,9,10,
// 	// }

// 	// fmt.Println(array)


// 	// fmt.Println(y[0])
// 	// fmt.Println(y[1])
// 	// fmt.Println(y[2])
// 	// fmt.Println(y[3])
// 	// fmt.Println(y[4])

// 	// primes := [6]int{2,3,5,7,11,13}

// 	// var s []int = primes[0:6]
// 	// fmt.Println(s)

// 	// s = append(s, 17, 19, 21, 23, 25)
// 	// fmt.Println(s)	

// 	// for i := 0; i < 10; i++ {
// 	// 	sum += i
// 	// 	fmt.Println(sum)
// 	// }
// 	// fmt.Println(sum)

// 	// if nilai > 90 {
// 	// 	fmt.Println("Nilai A")
// 	// } else if nilai > 80 {
// 	// 	fmt.Println("Nilai B")
// 	// } else if nilai > 70 {
// 	// 	fmt.Println("Nilai C")
// 	// } else {
// 	// 	fmt.Println("Nilai D")
// 	// }

// 	// if x{
// 	// 	fmt.Println("x is true")
// 	// } else {
// 	// 	fmt.Println("x is false")
// 	// }

// 	// fmt.Println("Tipe data sebelum dicasting: \n",dataFloat)
// 	// fmt.Printf("Tipe data setelah dicasting: \n%d\n", parseToInt) // Menambahkan %d untuk mencetak integer


// 	// isName := "John"
// 	// isNameModified := strings.ToUpper(isName)


// 	// fmt.Println(isNameModified)
// 	// fmt.Println(isName)


// 	// isNumber := "100"
// 	// parseToInt, err := strconv.Atoi(isNumber)

// 	// if err == nil {
// 	// 	fmt.Println(parseToInt)
// 	// } else {
// 	// 	fmt.Println(err)
// 	// }


// 	SayHai.SayHai("John", "10.00", "11.00", "12.00") // Pastikan SayHai didefinisikan di say.go
// 	SayHai.SayHaiv2("John", []string{"10.00", "11.00", "12.00"}) // Pastikan SayHai didefinisikan di say.go
// }


package main

import (
	"context"
	"fmt"
  
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
  )
  
  func main0() {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://fayruzaldi0:rXmAOOoKSdcMCHYz@learning.bdmmp.mongodb.net/?retryWrites=true&w=majority&appName=Learning").SetServerAPIOptions(serverAPI)
  
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
	  panic(err)
	}
  
	defer func() {
	  if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	  }
	}()
  
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
	  panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
  }
  