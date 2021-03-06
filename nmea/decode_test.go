package nmea

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

const examples = `
$GPGLL,4717.11634,N,00833.91297,E,124923.00,A,A*6E
$GPGLL,,,,,124924.00,V,N*42
$GPGLL,,,,,,V,N*64
$EIGBQ,RMC*28
$EIGLQ,RMC*26
$EIGNQ,RMC*24
$EIGPQ,RMC*3A
$GPDTM,W84,,0.0,N,0.0,E,0.0,W84*6F
$GPDTM,999,,0.08,N,0.07,E,-47.7,W84*1B
$GPGBS,235503.00,1.6,1.4,3.2,,,,,,*40
$GPGBS,235458.00,1.4,1.3,3.1,03,,-21.4,3.8,1,0*5A
$GPGGA,092725.00,4717.11399,N,00833.91590,E,1,08,1.01,499.6,M,48.0,M,,*5B
$GPGLL,4717.11364,N,00833.91565,E,092321.00,A,A*60
$GNGNS,103600.01,5114.51176,N,00012.29380,W,ANNN,07,1.18,111.5,45.6,,,V*00
$GNGNS,122310.2,3722.425671,N,12258.856215,W,DAAA,14,0.9,1005.543,6.5,,,V*0B
$GPGNS,122310.2,,,,,,07,,,,5.2,23,V*07
$GNGRS,104148.00,1,2.6,2.2,-1.6,-1.1,-1.7,-1.5,5.8,1.7,,,,,1,1*52
$GNGRS,104148.00,1,,0.0,2.5,0.0,,2.8,,,,,,,1,5*51
$GPGSA,A,3,23,29,07,08,09,18,26,28,,,,,1.94,1.18,1.54,1*10
$GPGST,082356.00,1.8,,,,1.7,1.3,2.2*7E
$GPGSV,3,1,09,09,,,17,10,,,40,12,,,49,13,,,35,1*6F
$GPGSV,3,2,09,15,,,44,17,,,45,19,,,44,24,,,50,1*64
$GPGSV,3,3,09,25,,,40,1*6E
$GPGSV,1,1,03,12,,,42,24,,,47,32,,,37,5*66
$GAGSV,1,1,00,2*76
$GPRMC,083559.00,A,4717.11437,N,00833.91522,E,0.004,77.52,091202,,,A,V*2D
$GPTXT,01,01,02,u-blox      ag -  www.u-blox.com*50
$GPTXT,01,01,02,ANTARIS      ATR0620  HW  00000040*47
$GPVLW,,N,,N,15.8,N,1.2,N*65
$GPVTG,77.52,T,,M,0.004,N,0.008,K,A*06
$GPZDA,082710.00,16,09,2002,00,00*64
$PUBX,41,1,0007,0003,19200,0*25
$PUBX,00,081350.00,4717.113210,N,00833.915187,E,546.589,G3,2.1,2.0,0.007,77.52,0.007,,0.92,1.19,0.77,9,0,0*5F
$PUBX,40,GLL,1,0,0,0,0,0*5D
$PUBX,03,11,23,-,,,45,010,29,-,,,46,013,07,-,,,42,015,08,U,067,31,42,025,10,U,195,33,46,026,18,U,326,08,39,026,17,-,,,32,015,26,U,306,66,48,025,27,U,073,10,36,026,28,U,089,61,46,024,15,-,,,39,014*0D
$PUBX,04,073731.00,091202,113851.00,1196,15D,1930035,-2660.664,43,*5D
`

func TestDecode(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(examples))
	for l := 1; s.Scan(); l++ {
		if len(s.Bytes()) == 0 {
			continue
		}
		msg, err := Decode(s.Bytes())
		if err != nil {
			t.Error(l, s.Text(), ":", err)
		}
		fmt.Println(msg)
	}
	if err := s.Err(); err != nil {
		t.Error(err)
	}

}
