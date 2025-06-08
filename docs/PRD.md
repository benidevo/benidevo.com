# Benjamin Idewor Portfolio Website

## Product Requirements Document (PRD)

---

## 1. Executive Summary

**Project Overview**: A modern, performance-focused portfolio website showcasing Benjamin Idewor's expertise in distributed systems, backend engineering, and systems programming.

**Technical Architecture**: Go monolith with HTMX interactivity, using GitHub repositories as the primary data source. Deployed on Google Cloud Run for optimal performance and cost efficiency.

**Primary Goal**: Establish a professional online presence that demonstrates technical depth and engineering excellence to the developer community.

---

## 2. Product Vision & Objectives

### Vision Statement

Create a developer-first portfolio that showcases complex system engineering projects through clear technical storytelling, demonstrating both breadth of technologies and depth of implementation.

### Primary Objectives

- **Technical Credibility**: Demonstrate expertise in Go, Java, Python, and distributed systems
- **Project Showcase**: Highlight 4-5 key personal projects with architectural depth
- **Performance Excellence**: Sub-second load times reflecting backend engineering standards
- **Community Engagement**: Enable easy discovery and exploration of open-source contributions

### Success Metrics

- Page load time < 800ms (P95)
- Technical project engagement (GitHub stars, forks)
- Professional network growth
- Community recognition

---

## 3. Target Audience

### Primary Audience: **Senior Backend Engineers & Tech Leaders**

- **Demographics**: 5+ years experience, distributed systems focus
- **Goals**: Discovering quality engineering work, potential collaboration
- **Behavior**: Values technical depth over visual flair, appreciates clean architecture

### Secondary Audience: **Open Source Community**

- **Demographics**: Developers interested in systems programming
- **Goals**: Learning from implementations, contributing to projects
- **Behavior**: Focused on code quality, documentation, and reusability

### Tertiary Audience: **Technical Recruiters & CTOs**

- **Demographics**: Non-technical but evaluation-focused
- **Goals**: Assessing technical capability and project impact
- **Behavior**: Looks for clear project outcomes, technology diversity

---

## 4. Design System & Visual Identity

### Color Palette

**Primary**: Deep Navy (`#0F172A`) - Technical sophistication, reliability
**Secondary**: Electric Blue (`#3B82F6`) - Innovation, technical precision
**Accent**: Emerald Green (`#10B981`) - Success states, growth, active elements
**Neutral**: Warm Gray (`#64748B`) - Text, backgrounds
**Background**: Off-white (`#FAFAFA`) - Clean, professional base

**Rationale**: Professional yet modern palette that works well in both light and dark contexts, with sufficient contrast for accessibility.

### Typography

**Primary**: Inter (Google Fonts)

- Headers: Inter 600 (Semi-bold)
- Body: Inter 400 (Regular)
- Code: JetBrains Mono 400

**Rationale**: Inter provides excellent readability at all sizes and has a technical, modern aesthetic. JetBrains Mono for code blocks ensures optimal developer experience.

### Visual Hierarchy

- **H1**: 3.5rem, Inter 600, Deep Navy
- **H2**: 2.5rem, Inter 600, Deep Navy
- **H3**: 1.875rem, Inter 600, Deep Navy
- **Body**: 1rem, Inter 400, Warm Gray
- **Code**: 0.875rem, JetBrains Mono, Electric Blue background

---

## 5. Content Strategy & Information Architecture

### Site Structure

```
/                    # Landing page with hero + featured projects
/projects           # Complete project showcase
/projects/{slug}    # Individual project deep-dive
/about             # Professional background + skills
/blog              # Technical writing (future)
/contact           # Simple contact form
```

### Featured Projects (GitHub Repository Mapping)

#### **1. Ascentio - Job Prospecting Platform**

- **Repository**: `ascentio` (Go, 4 stars)
- **Focus**: Full-stack Go application, API design, job matching algorithms
- **Key Metrics**: Multi-stage matching pipeline, automated cover letter generation
- **Technical Highlights**: Clean architecture, concurrent processing, external API integration

#### **2. Distributed Redis Implementation**

- **Repository**: `redis` (Go)
- **Focus**: Systems programming, protocol implementation, data structures
- **Key Metrics**: RESP protocol compliance, core command support
- **Technical Highlights**: Concurrent client handling, memory management, protocol parsing

#### **3. Event-Driven Order Fulfillment**

- **Repository**: `order-fulfillment` (Java, 1 star)
- **Focus**: Event sourcing, CQRS patterns, eventual consistency
- **Key Metrics**: Event replay capabilities, command/query separation
- **Technical Highlights**: Domain events, aggregate design, consistency guarantees

#### **4. Microservices Weather System**

- **Repository**: `weather-forecast-system` (Java/Spring Boot)
- **Focus**: Microservices architecture, service communication, data aggregation
- **Key Metrics**: Multiple service coordination, external API integration
- **Technical Highlights**: Service discovery, fault tolerance, data pipeline

### Content Sourcing Strategy

- **Project Data**: GitHub API (repository metadata, README parsing, release information)
- **Technical Details**: Structured markdown files in dedicated content repository
- **Code Examples**: Direct integration with GitHub repository file contents
- **Architecture Diagrams**: SVG files stored in repository, rendered inline

---

## 6. Technical Architecture

### Technology Stack

- **Backend**: Go 1.21+ (Gorilla Mux, html/template)
- **Frontend Enhancement**: HTMX 1.9+, Alpine.js 3.13+, Tailwind CSS 3.3+
- **Visual Effects**: Particles.js for subtle background animation
- **Data Source**: GitHub REST API v4
- **Deployment**: Google Cloud Run (containerized)
- **CDN**: Cloud Storage + Cloud CDN for static assets

### Data Architecture

```
GitHub Repository Structure:
portfolio-content/
├── projects/
│   ├── ascentio.md              # Frontmatter + detailed content
│   ├── redis-implementation.md
│   ├── order-fulfillment.md
│   └── weather-system.md
├── data/
│   ├── skills.json             # Technical skills matrix
│   ├── experience.json         # Professional timeline
│   └── featured.json           # Homepage configuration
└── assets/
    ├── diagrams/               # Architecture SVGs
    └── screenshots/            # Project visuals
```

### Performance Requirements

- **Cold Start**: < 1 second (Go binary optimization)
- **Page Load**: < 800ms (P95)
- **API Calls**: Cached responses (15-minute TTL)
- **Bundle Size**: < 100KB total (minimal JavaScript)
- **Memory Usage**: < 50MB container RAM

---

## 7. User Experience & Interface Design

### Landing Page Experience

**Hero Section**:

- Professional headshot (Yes - builds trust and personal connection)
- Clear value proposition: "Software Engineer specializing in Distributed Systems"
- Location: Berlin, Germany
- Subtle particles.js animation (dark blue dots, slow movement)

**Featured Projects Grid**:

- 2x2 card layout on desktop, single column on mobile
- Each card: Project title, brief description, tech stack badges, GitHub link
- Hover effects: Subtle elevation, tech stack highlight
- HTMX-powered: Click to expand project details inline

**Skills Overview**:

- Visual representation of technology proficiency
- Categories: Languages, Frameworks, Infrastructure, Architecture Patterns
- Interactive hover states showing experience level

### Project Detail Experience

**Architecture-First Approach**:

- Lead with system architecture diagram
- Problem statement → Solution approach → Implementation highlights → Results
- Tabbed interface: Overview, Architecture, Implementation, Results
- Code snippets with syntax highlighting
- Performance metrics and impact statements

**Navigation**:

- Sticky header with smooth scroll navigation
- Breadcrumb trails for deep project pages
- "Next Project" recommendations
- GitHub repository quick-access throughout

### Responsive Design

- **Desktop**: Full feature set, optimized for technical detail consumption
- **Tablet**: Condensed layouts, touch-friendly interactions
- **Mobile**: Simplified navigation, focus on content consumption
- **Accessibility**: WCAG 2.1 AA compliance, keyboard navigation, screen reader support

---

## 8. Interactive Features & Functionality

### HTMX Enhancement Points

- **Project Filtering**: Filter by technology, project type, or year
- **Inline Project Details**: Expand project information without page navigation
- **Contact Form**: Async submission with feedback
- **GitHub Integration**: Live repository statistics and recent activity
- **Search**: Real-time project and content search

### Alpine.js Micro-Interactions

- **Skill Cards**: Hover to reveal proficiency details and project usage
- **Project Cards**: Smooth state transitions and progressive disclosure
- **Navigation**: Mobile menu toggle, smooth scroll behaviors
- **Form Validation**: Real-time input validation and user feedback

### GitHub Integration Features

- **Live Repository Data**: Stars, forks, last updated, primary language
- **Recent Activity**: Latest commits or releases from featured projects
- **Code Quality Indicators**: Repository health scores where available
- **Direct Links**: Seamless navigation to source code and documentation

---

## 9. Content Management & Maintenance

### Content Update Workflow

1. **Project Updates**: Edit markdown files in content repository
2. **New Projects**: Create new markdown file with frontmatter metadata
3. **Technical Details**: Update architecture diagrams or code examples
4. **Deployment**: Git push triggers automatic cache invalidation

### Content Quality Standards

- **Technical Accuracy**: All code examples tested and validated
- **Architecture Clarity**: Diagrams updated to reflect current implementation
- **Impact Metrics**: Quantified results where measurable
- **Accessibility**: Alt text for all images, proper heading hierarchy

### Automated Maintenance

- **GitHub API Caching**: 15-minute cache with graceful degradation
- **Repository Monitoring**: Automatic updates for stars, forks, recent activity
- **Performance Monitoring**: Automated alerts for load time regressions
- **Security Updates**: Dependabot integration for Go dependencies

---

## 10. Technical Implementation Notes

### Go Application Structure

```
portfolio/
├── main.go                 # Application entry point
├── handlers/               # HTTP request handlers
├── services/              # GitHub API integration, caching
├── models/                # Data structures for projects, skills
├── templates/             # HTML templates with HTMX integration
├── static/                # CSS, JS, images
└── config/                # Environment configuration
```

### Cloud Run Deployment Strategy

- **Container Optimization**: Multi-stage build, minimal base image
- **Environment Variables**: GitHub token, repository configuration
- **Health Checks**: Custom endpoints for availability monitoring
- **Scaling**: Min 0, max 10 instances with CPU-based scaling
- **Cost Optimization**: 0.25 vCPU, 512Mi memory allocation

### Security Considerations

- **GitHub Token**: Minimum required permissions (public repository read)
- **Rate Limiting**: Respect GitHub API limits with exponential backoff
- **Content Security Policy**: Strict CSP headers for XSS prevention
- **HTTPS Enforcement**: Automatic redirects and security headers

---

## 11. Launch Strategy & Success Metrics

### Soft Launch Phase (Week 1-2)

- **Internal Testing**: Performance validation, content accuracy review
- **Peer Review**: Feedback from 3-5 senior engineers in network
- **SEO Foundation**: Meta tags, structured data, sitemap generation

### Public Launch Phase (Week 3-4)

- **Community Sharing**: LinkedIn, Twitter, relevant Discord/Slack communities
- **GitHub Profile Update**: Link portfolio from profile README
- **Project README Updates**: Cross-link from featured project repositories

### Growth Metrics (3 months)

- **Traffic**: 500+ unique visitors monthly
- **Engagement**: 60%+ project detail view rate
- **Technical Community**: 50+ GitHub profile views weekly
- **Professional Network**: 25+ new LinkedIn connections

### Long-term Evolution

- **Technical Blog**: Monthly posts on systems engineering topics
- **Open Source Showcase**: Additional project highlights as portfolio grows
- **Community Contributions**: Speaking engagements, conference talks
- **Performance Optimization**: Continuous improvement based on real user metrics

---

## 12. Risk Assessment & Mitigation

### Technical Risks

- **GitHub API Rate Limits**: Mitigation through aggressive caching and fallback content
- **Repository Availability**: Graceful degradation with cached content
- **Performance Regression**: Automated monitoring with alerting thresholds

### Content Risks

- **Outdated Information**: Automated checks for repository metadata freshness
- **Technical Accuracy**: Peer review process for complex technical explanations
- **Professional Representation**: Content review against professional standards

### Business Continuity

- **Platform Independence**: Content structure supports easy migration
- **Backup Strategy**: Git-based content with multiple hosting options
- **Cost Management**: Cloud Run's pay-per-use model with spending alerts

---

*Document Version: 1.0*
*Last Updated: June 8, 2025*
*Next Review: July 8, 2025*
