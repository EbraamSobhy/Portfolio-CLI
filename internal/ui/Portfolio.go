package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/EbraamSobhy/Portfolio-CLI/internal/styles"
)

type sectionItem struct {
	key         string
	title       string
	subtitle    string
	description string
	category    string
}

func (i sectionItem) Title() string       { return i.title }
func (i sectionItem) Description() string { return i.subtitle }
func (i sectionItem) FilterValue() string { return i.title }

type menuModel struct {
	list     list.Model
	width    int
	height   int
	quantity int
}

func NewMenuModel() tea.Model {
	items := []list.Item{
		// Featured
		sectionItem{
			key: "Ebraam Sobhy",
			title: "Ebraam Sobhy",
			subtitle: "Frontend Developer | Mobile Developer | Software Development Instructor",
			description: "Experienced web and mobile developer and instructor specializing in frontend and Mobile (React Native) technologies, passionate about creating and teaching and practical skill-building in the ever-evolving field of web and mobile development",
			category: "Hi, I am",
		},
		// Portfolio
		sectionItem{
			key: "My Services",
			title: "My Services",
			subtitle: "Frontend Development | Mobile App Development | Software Development Instructor",
			description: "I specialize in Frontend Web Development, creating visually appealing, responsive, and user-friendly websites. Utilizing modern technologies, I build fast, accessible interfaces that feel great to use.\n\nAs a Software Development Instructor, I teach essential web development skills, combining theory with hands-on practice. My goal is to empower aspiring developers with the knowledge and confidence needed to excel in the dynamic field of software development.\n\nWith expertise in Mobile App Development using React Native, I create seamless and high-performing cross-platform mobile applications. Proficient in JavaScript and React Native, I ensure your app delivers an excellent user experience and meets all functional and design requirements.",
			category: "Portfolio",
		},
		sectionItem{
			key: "My Skills",
			title: "My Skills",
			subtitle: "Web Development | Mobile Development | Information Systems | Software Developer Instructor",
			description: "HTML, CSS, JavaScript, TypeScript , Tailwind CSS, React JS, Next JS, React Native, Node.js (Express.js), Python, FastAPI, Golang, Cobra CLI,\n\nMongoDB, MySQL , Firebase, Supabase, Docker, Linux, Bash, Web Scraping (BeautifulSoup), CI/CD pipeline (GitHub Actions), Git & GitHub",
			category: "Portfolio",
		},
		sectionItem{
			key: "Top Projects",
			title: "Top Projects",
			subtitle: "This is my top web development and CLI apps\n",
			description: "Cyber Security\n\nProject Link: [https://github.com/EbraamSobhy/cyber-security]\n\n\nGammal Tech Final Exam\n\nProject Link: [https://gammal-tech-final-exam.vercel.app/]\n\n\nDevSheets\n\nProject Link: [https://devsheets-cli-app.netlify.app/]\n\n\nClash of Patrols\n\nProject Link: [https://www.linkedin.com/feed/update/urn:li:activity:7356932987407544321/]\n\n\nDevPath\n\nProject Link: [https://devpathapp.netlify.app/]\n\n\n",
			category: "Portfolio",
		},
		sectionItem{
			key: "Top Mobile Projects",
			title: "Top Mobile Projects",
			subtitle: "This is my top mobile development apps\n",
			description: "My Portfolio\n\nProject Link: [https://portfolio-three-js-ruddy.vercel.app/assets/My%20Portfolio.mp4]\n\n\nLogin and SignUp\n\nProject Link: [https://portfolio-three-js-ruddy.vercel.app/assets/Login-signup-nodejs.mp4]\n\n\nNote Taking\n\nProject Link: [https://portfolio-three-js-ruddy.vercel.app/assets/Note-Taking-App.mp4]\n\n\nDevPath\n\nProject Link: [https://portfolio-three-js-ruddy.vercel.app/assets/DevPath.mp4]\n\n\n",
			category: "Portfolio",
		},

		sectionItem{
			key: "Social Media",
			title: "Social Media",
			subtitle: "This is my social media links\n",
			description: "Gmail\n\nLink: [ebraam.sobhy2003@gmail.com]\n\n\nLinkedin\n\nLink: [https://www.linkedin.com/in/ebraam-sobhy-255444274/]\n\n\nGithub\n\nLink: [https://github.com/EbraamSobhy]",
			category: "Portfolio",
		},
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Background(lipgloss.Color(styles.ColorPrimary)).
		Foreground(lipgloss.Color(styles.ColorTextOnSel)).
		Padding(0, 1)
	delegate.Styles.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorMuted)).
		Padding(0, 1)

	l := list.New(items, delegate, 0, 0)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	// Start at index 0 to highlight [Ebraam Sobhy]
	l.Select(0)

	return menuModel{
		list: l,
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "+":
			m.quantity++
		case "-":
			if m.quantity > 0 {
				m.quantity--
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m menuModel) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}

	topBar := m.renderTopBar()
	tabs := m.renderTabs()
	body := m.renderBody()

	footer := m.renderFooter()

	content := lipgloss.JoinVertical(lipgloss.Left,
		topBar,
		tabs,
		body,
		footer,
	)

	// Wrap in a style that ensures it takes the full width and has some margin
	return lipgloss.NewStyle().Padding(1, 2).Render(content)
}

func (m menuModel) renderTopBar() string {
	left := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#0E8ED5")).
		Padding(0, 1).
		Render("Ebraam Sobhy")

	leftSide := lipgloss.JoinHorizontal(lipgloss.Top, left, " ")

	rightSide := lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorSecondary)).
		Bold(true).
		Render("Software Developer")

	gapWidth := m.width - lipgloss.Width(leftSide) - lipgloss.Width(rightSide) - 4
	if gapWidth < 0 {
		gapWidth = 0
	}
	gap := lipgloss.NewStyle().Width(gapWidth).Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, leftSide, gap, rightSide)
}

func (m menuModel) renderTabs() string {
	thinBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	tabStyle := lipgloss.NewStyle().
		Border(thinBorder, false, true, false, true).
		BorderForeground(lipgloss.Color(styles.ColorTextOnSel)).
		Padding(0, 1).
		Align(lipgloss.Center)

	activeTabStyle := tabStyle.Copy().
		Foreground(lipgloss.Color(styles.ColorForeground)).
		Bold(true)

	inactiveTabStyle := tabStyle.Copy().
		Foreground(lipgloss.Color(styles.ColorForeground))

	tabs := lipgloss.JoinHorizontal(lipgloss.Top,
			activeTabStyle.Render("About Me"),
		inactiveTabStyle.Render("Services"),
		inactiveTabStyle.Render("Skills"),
		inactiveTabStyle.Render("Top Projects"),
	)

	return lipgloss.NewStyle().
		Border(thinBorder, false, false, true, false).
		BorderForeground(lipgloss.Color(styles.ColorTextOnSel)).
		Width(m.width - 12).
		Align(lipgloss.Center).
		Render(tabs)
}

func (m menuModel) renderBody() string {
	sidebarWidth := 30
	contentWidth := m.width - sidebarWidth - 4

	sidebar := m.renderSidebar(sidebarWidth)
	content := m.renderContent(contentWidth)

	bodyHeight := m.height - 12
	if bodyHeight < 10 {
		bodyHeight = 10
	}

	return lipgloss.NewStyle().Height(bodyHeight).Render(lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content))
}

func (m menuModel) renderSidebar(width int) string {
	var lines []string

	// Featured Header
	lines = append(lines, "")
	lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color(styles.ColorMuted)).Render(" ~ Hi, I am ~"))

	// Featured items
	for i, item := range m.list.Items() {
		it := item.(sectionItem)
		if it.category == "Hi, I am" {
			lines = append(lines, m.renderMenuItem(it, i == m.list.Index()))
		}
	}

	// Portfolio Header
	lines = append(lines, "")
	lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color(styles.ColorMuted)).Render(" ~ Portfolio ~"))

	// Portfolio items
	for i, item := range m.list.Items() {
		it := item.(sectionItem)
		if it.category == "Portfolio" {
			lines = append(lines, m.renderMenuItem(it, i == m.list.Index()))
		}
	}

	return lipgloss.NewStyle().Width(width).Padding(1, 2).Render(strings.Join(lines, "\n"))
}

func (m menuModel) renderMenuItem(it sectionItem, selected bool) string {
	if selected {
		return lipgloss.NewStyle().
			Background(lipgloss.Color(styles.ColorPrimary)).
			Foreground(lipgloss.Color(styles.ColorTextOnSel)).
			Padding(0, 1).
			Render(it.title)
	}
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorMuted)).
		Padding(0, 1).
		Render(it.title)
}

func (m menuModel) renderContent(width int) string {
	item, ok := m.list.SelectedItem().(sectionItem)
	if !ok {
		return ""
	}

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorForeground)).
		Bold(true).
		Render(item.title)

	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorTextOnSel)).
		Render(item.subtitle)

	description := lipgloss.NewStyle().
		Foreground(lipgloss.Color(styles.ColorTextOnSel)).
		Width(width - 10).
		Render(item.description)



	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		subtitle,
		"",
		description,
		"",

	)

	return lipgloss.NewStyle().Padding(2, 4).Render(content)
}



func (m menuModel) renderFooter() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(styles.ColorMuted))
	thinBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	Portfolio := style.Render("↑/↓ Portfolio")
	quit := style.Render("q quit")

	footer := lipgloss.JoinHorizontal(lipgloss.Top,
		Portfolio, "   ", quit,
	)

	return lipgloss.NewStyle().
		Border(thinBorder, true, false, false, false).
		BorderForeground(lipgloss.Color(styles.ColorTextOnSel)).
		Width(m.width - 6).
		Align(lipgloss.Center).
		Render(footer)
}
