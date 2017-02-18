package main

import (
	"github.com/andlabs/ui"
)

var window *ui.Window

func main() {
	ui.Main(func() {
		window = ui.NewWindow("libui Control Gallery", 400, 300, true)
		window.SetMargined(true)

		tab := ui.NewTab()
		tab.Append("Basic Controls", makeBasicControlsPage())
		tab.SetMargined(0, true)
		tab.Append("Numbers and Lists", makeNumbersPage())
		tab.SetMargined(1, true)
		tab.Append("Data Choosers", makeDataChoosersPage())
		tab.SetMargined(2, true)
		window.SetChild(tab)

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})
}

func makeNumbersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	var spinbox *ui.Spinbox
	var slider *ui.Slider
	var pbar *ui.ProgressBar

	spinbox = ui.NewSpinbox(0, 100)
	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	vbox.Append(spinbox, false)

	slider = ui.NewSlider(0, 100)
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(slider, false)

	pbar = ui.NewProgressBar()
	vbox.Append(pbar, false)

	ip := ui.NewProgressBar()
	ip.SetValue(0)
	vbox.Append(ip, false)

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	cbox.SetSelected(0)
	vbox.Append(cbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}

func makeBasicControlsPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	hbox.Append(ui.NewButton("Button"), false)
	hbox.Append(ui.NewCheckbox("Checkbox"), false)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)
	vbox.Append(ui.NewHorizontalSeparator(), false)
	vbox.Append(hbox, false)

	group := ui.NewGroup("Entries")
	group.SetMargined(true)
	vbox.Append(group, true)

	entryForm := ui.NewEntry()
	group.SetChild(entryForm)

	return vbox
}

func makeDataChoosersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox1 := ui.NewVerticalBox()
	vbox1.SetPadded(true)
	hbox.Append(vbox1, false)

	vbox1.Append(ui.NewDatePicker(), false)
	vbox1.Append(ui.NewTimePicker(), false)
	vbox1.Append(ui.NewDateTimePicker(), false)

	vbox2 := ui.NewVerticalBox()
	vbox2.SetPadded(true)
	hbox.Append(vbox2, true)

	hbox1 := ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	vbox2.Append(hbox1, false)

	entry1 := ui.NewEntry()
	entry1.SetReadOnly(true)
	entry1.SetText("entry1")
	hbox1.Append(entry1, false)

	button1 := ui.NewButton("Open File")
	button1.OnClicked(func(b *ui.Button) {
		fileName := ui.OpenFile(window)
		if fileName == "" {
			entry1.SetText("(cancelled)")
			return
		}
		entry1.SetText(fileName)
	})
	hbox1.Append(button1, false)

	hbox2 := ui.NewHorizontalBox()
	hbox2.SetPadded(true)
	vbox2.Append(hbox2, false)

	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	entry2.SetText("entry2")
	hbox2.Append(entry2, false)

	button2 := ui.NewButton("Save File")
	button2.OnClicked(func(b *ui.Button) {
		fileName := ui.SaveFile(window)
		if fileName == "" {
			entry2.SetText("(cancelled)")
			return
		}
		entry2.SetText(fileName)
	})
	hbox2.Append(button2, false)

	hbox3 := ui.NewHorizontalBox()
	hbox3.SetPadded(true)
	vbox2.Append(hbox3, false)

	button3 := ui.NewButton("Message Box")
	button3.OnClicked(func(b *ui.Button) {
		ui.MsgBox(window, "This is a normal message box.", "More detailed information can be shown here.")
	})
	hbox3.Append(button3, false)

	button4 := ui.NewButton("Error Box")
	button4.OnClicked(func(b *ui.Button) {
		ui.MsgBox(window, "This message box describes an error.", "More detailed information can be shown here.")
	})
	hbox3.Append(button4, false)

	return hbox
}
