package SRP_net

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/zexy-swami/SRP/SRP_CLI/internal/SRP"
	"github.com/zexy-swami/SRP/SRP_CLI/pkg/parser"
)

func StartClientConnection() error {
	conn, err := net.Dial("tcp", getAddress())
	if err != nil {
		return err
	}
	defer conn.Close()

	srpID := parser.GetDataFromConfig("srp_id")
	conn.Write([]byte(srpID + "\n"))

	authenticationResult, _ := bufio.NewReader(conn).ReadString('\n')
	authenticationResult = strings.Trim(authenticationResult, "\n")
	if strings.Contains(authenticationResult, "not") {
		return errors.New("error occurred while authorizing with srp id")
	}

	pfIndex, _ := strconv.Atoi(authenticationResult)
	clientProver := SRP.NewProver(pfIndex, srpID)

	publicAValue := clientProver.GenerateClientPublicValue()
	conn.Write([]byte(publicAValue.String() + "\n"))

	publicServerValueAndSalt, _ := bufio.NewReader(conn).ReadString('\n')
	publicServerValueAndSalt = strings.Trim(publicServerValueAndSalt, "\n")
	publicServerValueAndSaltAsSlc := strings.Split(publicServerValueAndSalt, "<:>")

	clientMValue := clientProver.GenerateMValue(publicServerValueAndSaltAsSlc[0], publicServerValueAndSaltAsSlc[1])
	conn.Write([]byte(clientMValue.String() + "\n"))

	testValueFromServer, _ := bufio.NewReader(conn).ReadString('\n')
	testValueFromServer = strings.Trim(testValueFromServer, "\n")
	if strings.Contains(testValueFromServer, "wrong") {
		return errors.New("error occurred while verification")
	}
	if isEq := clientProver.CompareZValues(testValueFromServer); !isEq {
		return errors.New("error occurred while comparing z values")
	}

	fmt.Println("verification passed successfully")
	conn.Write([]byte("fin.\n"))
	
	return nil
}

func getAddress() string {
	host := parser.GetDataFromConfig("host")
	port := parser.GetDataFromConfig("port")
	return host + ":" + port
}
