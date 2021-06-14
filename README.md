# Minimal User Interface

Functionally Minimal HTML Generation

* generates tags with indention automatically
* can embed raw HTML strings as well
* extremely minimal html output for fast downloading
* suitable for internal or hastily built interfaces

## Example

	menu := fui.NavItems{
		fui.NavItem{"Home", "/"},
		fui.NavItem{"Test", "/test"},
	}
	
	html := fui.Page("Title", "cosmo",
		fui.Navbar(menu, "ACME Corp", r.URL.Path) + "\n" +
		fui.Tag("H1", nil, "Hello World"))
