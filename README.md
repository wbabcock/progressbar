# Terminal Progress Bar

### Simple terminal progress bar in Go

This is a super simple, but configurable progress bar for your Go CLI Project.

```go
p := progressbar.NewProgressBar()
p.Update(10) // will display a 10% progress bar
```
This is all that is needed for a progress bar.

You can also configure to fit your needs.

```go
p := progressbar.NewProgressBar().
		WithLength(80).
		WithIndicator(indicator.Block).
		WithColor(colors.Green)

p.Update(10) // will display a 10% progress bar
```

You can adjust the length of the bar, what type of indicator, and the color.

Indicator Types: 
- Hash (default)
- Star       
- Underscore 
- Minus      
- Equals     
- Slash      
- Block      
- Square     

The colors are the basic 8 terminal colors: Black, Red, Blue, Yellow, Green, Magenta, Cyan, White.