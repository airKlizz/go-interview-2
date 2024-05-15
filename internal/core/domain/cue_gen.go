// Code generated by gocode.Generate; DO NOT EDIT.

package domain

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
	_ "cuelang.org/go/pkg"
)

var cuegenvalEvent = cuegenMake("Event", &Event{})

// Validate validates x.
func (x *Event) Validate() error {
	return cuegenCodec.Validate(cuegenvalEvent, x)
}

var cuegenCodec, cuegenInstance_, cuegenValue = func() (*gocodec.Codec, *cue.Instance, cue.Value) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0], instances[0].Value()
}()

// Deprecated: cue.Instance is deprecated. Use cuegenValue instead.
var cuegenInstance = cuegenInstance_

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenValue.FieldByName(name, true)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 417 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xfft\x90\xefj\xd4@\x14\xc5\xef\u036e\xe0\\\xaao \f)\xd4\xfa\u014d\x88E\x06+\xf4\x8f\xf6\x9bH\x11\xfc \"\xc3\xf4n2\x9a\xcc,\x93I\x8b\xacA\xad\xd5\xd7\xf1\xd9|\x81F\x92\u01ad\xff\xfaq\xce9\xf3\xe3\x9es\xa3\xfb\x9a`\xd2}\x03\xec>\x01l}\x9c \xaeYWG\xed\f\xef\xeb\xa8{\x19'8=\xf4>b\x028}\xaec\x81k\x80\u05de\u0692k\xec\xce\x00\xe0V\xf7%A\xbc\xf9\xea\xb5i\xf8\xee\u0716\xe3\xcf3\xc0\xee\x14`\xb3\xfb<A\xbc~\xa9\x9f\x02&8}\xa6+\xeeA\xd3A$\x008\xc7\x1f\xfd!\x88x\xbbz\xef\xf8$\xf7\x8b\xe0\u07f2\x893\xeb\"\a\xa7\u02d9\xf1\x81gG\xbe\xd2\xd6\xcdL\u00c8(\x1aW\xe9P\x17\xba\xc4s\xfc\xbe\xd0\xe6\x9d\xceY\x9a\x86\x89\xd6\xf7|\u90d2K\x12\x87|\xa4\xa4\x94\x8du\xf1!\x89\x83\xc0\xec\u052f\xd7n\u0670Zy/\v\x1by\xe5\x1dh\xebFOn\xc8G\xdb\xf7\xb2\x8cZZ\x1fCK\x12/\xb8Z\xf4d)W\xa9\xc7\xdb\xf7\xb3,\x1b\xe2[\x0f\xb2\x8c\xc4n\xb0y\x11\x1d\u05f5\xfa\x9b\xf4\xe4\x98]TrsIb\xc7D\ub752\xa9w)\xb5\xf2\x83\xfcC\x9b\xcf\xff\x15M\xa1]\xceoL\xdf2%\xb1\x13\xf2z\xb8I\xec\r\xc6\xd0\xfeR\x14\xe3\x1a\x17\xab\x90\x10-\x89\xf6*\xe6I_\xf0?\u0321\xf8o\xccq\x88\x8bAV\xcc;rc\xd8F\x87\x9c\xa3\x92u\f\xd6\xe5$\xf6\xf9\xd8\x1aV2-\xfb=Rj\t\xe0g\x00\x00\x00\xff\xff$`\x1e#\x80\x02\x00\x00")
