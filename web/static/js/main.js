// Main JavaScript file for portfolio website
// Configurable particles.js and other interactions

// Particles.js configuration - easily customizable
const particlesConfig = {
  particles: {
    number: {
      value: 80,
      density: {
        enable: true,
        value_area: 800
      }
    },
    color: {
      value: "#3B82F6" // Uses CSS variable --color-secondary
    },
    shape: {
      type: "circle",
      stroke: {
        width: 0,
        color: "#000000"
      }
    },
    opacity: {
      value: 0.5,
      random: false,
      anim: {
        enable: false,
        speed: 1,
        opacity_min: 0.1,
        sync: false
      }
    },
    size: {
      value: 3,
      random: true,
      anim: {
        enable: false,
        speed: 40,
        size_min: 0.1,
        sync: false
      }
    },
    line_linked: {
      enable: true,
      distance: 150,
      color: "#3B82F6",
      opacity: 0.4,
      width: 1
    },
    move: {
      enable: true,
      speed: 2,
      direction: "none",
      random: false,
      straight: false,
      out_mode: "out",
      bounce: false,
      attract: {
        enable: false,
        rotateX: 600,
        rotateY: 1200
      }
    }
  },
  interactivity: {
    detect_on: "canvas",
    events: {
      onhover: {
        enable: true,
        mode: "repulse"
      },
      onclick: {
        enable: true,
        mode: "push"
      },
      resize: true
    },
    modes: {
      grab: {
        distance: 400,
        line_linked: {
          opacity: 1
        }
      },
      bubble: {
        distance: 400,
        size: 40,
        duration: 2,
        opacity: 8,
        speed: 3
      },
      repulse: {
        distance: 100,
        duration: 0.4
      },
      push: {
        particles_nb: 4
      },
      remove: {
        particles_nb: 2
      }
    }
  },
  retina_detect: true
};

// Initialize particles.js when DOM is ready
document.addEventListener('DOMContentLoaded', function() {
  // Initialize particles if the container exists
  if (document.getElementById('particles-js')) {
    particlesJS('particles-js', particlesConfig);
  }

  // Smooth scroll for anchor links
  document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
      e.preventDefault();
      const target = document.querySelector(this.getAttribute('href'));
      if (target) {
        target.scrollIntoView({
          behavior: 'smooth',
          block: 'start'
        });
      }
    });
  });

  // Animation on scroll (simple intersection observer)
  const observeElements = document.querySelectorAll('.animate-slide-up, .animate-fade-in');
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.style.opacity = '1';
        entry.target.style.transform = 'translateY(0)';
      }
    });
  }, {
    threshold: 0.1
  });

  observeElements.forEach(el => {
    el.style.opacity = '0';
    el.style.transform = 'translateY(20px)';
    el.style.transition = 'opacity 0.6s ease-out, transform 0.6s ease-out';
    observer.observe(el);
  });
});

// HTMX event handlers for enhanced interactivity
document.addEventListener('htmx:afterRequest', function(event) {
  if (event.target.id && event.target.id.startsWith('project-details-')) {
    const detailsEl = event.target;
    if (detailsEl.innerHTML.trim()) {
      detailsEl.classList.remove('hidden');
      detailsEl.style.opacity = '0';
      detailsEl.style.transform = 'translateY(-10px)';

      // Animate in
      requestAnimationFrame(() => {
        detailsEl.style.transition = 'opacity 0.3s ease-out, transform 0.3s ease-out';
        detailsEl.style.opacity = '1';
        detailsEl.style.transform = 'translateY(0)';
      });
    }
  }
});

// Utility function to update theme colors dynamically
function updateThemeColors(colors) {
  const root = document.documentElement;
  Object.entries(colors).forEach(([key, value]) => {
    root.style.setProperty(`--color-${key}`, value);
  });

  // Update particles color if they exist
  if (window.pJSDom && window.pJSDom[0] && window.pJSDom[0].pJS) {
    window.pJSDom[0].pJS.particles.color.value = colors.secondary || '#3B82F6';
    window.pJSDom[0].pJS.particles.line_linked.color = colors.secondary || '#3B82F6';
  }
}

// Theme toggle functionality
function toggleTheme() {
  const currentTheme = document.documentElement.getAttribute('data-theme');
  const newTheme = currentTheme === 'light' ? 'dark' : 'light';

  if (newTheme === 'light') {
    document.documentElement.setAttribute('data-theme', 'light');
  } else {
    document.documentElement.removeAttribute('data-theme');
  }

  localStorage.setItem('theme', newTheme);

  updateParticlesTheme(newTheme);
}

function updateParticlesTheme(theme) {
  if (window.pJSDom && window.pJSDom[0] && window.pJSDom[0].pJS) {
    const color = theme === 'light' ? '#3B82F6' : '#60A5FA';
    window.pJSDom[0].pJS.particles.color.value = color;
    window.pJSDom[0].pJS.particles.line_linked.color = color;
  }
}

// Initialize theme on page load
function initializeTheme() {
  const savedTheme = localStorage.getItem('theme');

  const themeToApply = savedTheme || 'dark';

  if (themeToApply === 'light') {
    document.documentElement.setAttribute('data-theme', 'light');
  } else {
    document.documentElement.removeAttribute('data-theme');
  }

  updateParticlesTheme(themeToApply);
}

initializeTheme();

// Export for global access
window.portfolioUtils = {
  updateThemeColors,
  particlesConfig
};

// Make theme toggle globally available
window.toggleTheme = toggleTheme;