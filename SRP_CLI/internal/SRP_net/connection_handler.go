package SRP_net

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/zexy-swami/SRP/SRP_CLI/internal/SRP"
)

func connectionHandler(connection net.Conn) {
	srpID, _ := bufio.NewReader(connection).ReadString('\n')
	srpID = strings.Trim(srpID, "\n")
	if authorized := SRP.CheckUserID(srpID); !authorized {
		connection.Write([]byte("Sorry, you're not authorized\n"))
		return
	}

	pfIndex := SRP.GetPrimeFieldIndex()
	pfIndexAsString := strconv.Itoa(pfIndex) + "\n"
	connection.Write([]byte(pfIndexAsString))

	serverVerifier := SRP.NewVerifier(pfIndex, srpID)
	publicClientValue, _ := bufio.NewReader(connection).ReadString('\n')
	publicClientValue = strings.Trim(publicClientValue, "\n")

	publicServerValue, usedSalt := serverVerifier.GenerateServerPublicValueAndSalt(publicClientValue)
	connection.Write([]byte(publicServerValue.String() + "<:>" + usedSalt.String() + "\n"))

	clientMValue, _ := bufio.NewReader(connection).ReadString('\n')
	clientMValue = strings.Trim(clientMValue, "\n")

	if isEq := serverVerifier.CompareMValues(clientMValue); !isEq {
		connection.Write([]byte("wrong check value\n"))
		fmt.Println("User wasn't verified correctly")
		return
	}
	fmt.Printf("User %s was verified correctly\n", srpID)

	serverTestValue := serverVerifier.GenerateZValue()
	connection.Write([]byte(serverTestValue.String() + "\n"))
}
