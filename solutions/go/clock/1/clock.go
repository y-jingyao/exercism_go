package clock

import "strconv"

// Define the Clock type here.

type Clock struct {
    mi int
}

func New(h, m int) Clock {
	//panic("Please implement the New function")
    day := 24 * 60
	total := h*60 + m
	total = total % day
	if total < 0 {
		total += day
	}
	return Clock{mi: total}
}

func (c Clock) Add(m int) Clock {
	//panic("Please implement the Add function")
    c.mi += m
    c.mi = c.mi % (24*60)
    if c.mi < 0 {
        c.mi += (24*60)
    }
    return c
}

func (c Clock) Subtract(m int) Clock {
	//panic("Please implement the Subtract function")
    c.mi -= m
    c.mi = c.mi % (24*60)
     if c.mi < 0 {
        c.mi += (24*60)
    }
    return c
}

func (c Clock) String() string {
	//panic("Please implement the String function")
    h := c.mi / 60
    if h == 24 {
        h = 0
    } 
    m := c.mi - 60*h
    strh := strconv.Itoa(h)
    if h<10 {
        strh = "0" + strh	
    }
    strm := strconv.Itoa(m)
    if m<10 {
        strm = "0" + strm
    }
    return strh + ":" + strm
}
