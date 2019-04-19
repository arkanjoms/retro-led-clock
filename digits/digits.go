package digits

type Number [5]string

const Separator string = "░"

var Zero = Number{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var One = Number{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var Two = Number{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var Three = Number{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var Four = Number{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var Five = Number{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}
var Six = Number{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var Seven = Number{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var Eigth = Number{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var Nine = Number{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var Numbers = [...]Number{
	Zero,
	One,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eigth,
	Nine,
}
