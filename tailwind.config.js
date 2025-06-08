/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./web/templates/**/*.{html,js}",
    "./web/static/js/**/*.js"
  ],
  theme: {
    extend: {
      colors: {
        // Custom color palette - easily configurable
        'primary': 'var(--color-primary, #0F172A)',       // Deep Navy
        'secondary': 'var(--color-secondary, #3B82F6)',   // Electric Blue
        'accent': 'var(--color-accent, #10B981)',         // Emerald Green
        'neutral': 'var(--color-neutral, #64748B)',       // Warm Gray
        'background': 'var(--color-background, #FAFAFA)', // Off-white
        'surface': 'var(--color-surface, #FFFFFF)',       // White for cards
      },
      fontFamily: {
        // Custom fonts - easily configurable
        'sans': ['var(--font-sans, Inter)', 'system-ui', 'sans-serif'],
        'mono': ['var(--font-mono, "JetBrains Mono")', 'monospace'],
      },
      fontSize: {
        // Custom font sizes following the design system
        'h1': ['3.5rem', { lineHeight: '1.2', fontWeight: '600' }],
        'h2': ['2.5rem', { lineHeight: '1.3', fontWeight: '600' }],
        'h3': ['1.875rem', { lineHeight: '1.4', fontWeight: '600' }],
        'body': ['1rem', { lineHeight: '1.6', fontWeight: '400' }],
        'code': ['0.875rem', { lineHeight: '1.5', fontWeight: '400' }],
      },
      animation: {
        'fade-in': 'fadeIn 0.3s ease-in-out',
        'slide-up': 'slideUp 0.3s ease-out',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
        slideUp: {
          '0%': { transform: 'translateY(10px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' },
        },
      },
    },
  },
  plugins: [],
}