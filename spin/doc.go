// Package spin provides a thread safe Spinner with various available character sets, prefix, suffix and color.
// The Spinner can be controlled using an Options pattern.
// Example:
//
//	s := spin.New(spin.BouncingBar, time.Millisecond * 10,
//	    spin.WithPrefix("SPINNING "),
//	    spin.WithSuffix("AFTER"),
//	    spin.WithColor(color.FgGreen),
//	    spin.WithFinalMsg("BYE!"))
//
//	s.Start()
//	time.Sleep(time.Second * 3)
//	s.Stop()
package spin

