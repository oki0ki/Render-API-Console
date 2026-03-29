/** @type {import('tailwindcss').Config} */
export default {
	darkMode: 'class',
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				sans: ['Inter', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'sans-serif']
			},
			borderRadius: {
				DEFAULT: '1rem',
				ui: '1rem',
				sm: '0.7rem'
			},
			colors: {
				gray: {
					50: '#f9f9f9',
					100: '#ececec',
					200: '#e3e3e3',
					300: '#cdcdcd',
					400: '#b4b4b4',
					500: '#9b9b9b',
					600: '#676767',
					700: '#4e4e4e',
					800: '#333',
					850: '#262626',
					900: '#171717',
					950: '#0d0d0d'
				},
				surface: {
					DEFAULT: '#ffffff',
					secondary: '#F9F9F9',
					tertiary: '#F2F2F2',
					dark: '#000000',
					'dark-secondary': '#191919',
					'dark-tertiary': '#212121',
					'dark-quaternary': '#2C2C2E'
				},
				border: {
					DEFAULT: '#EBEBEB',
					dark: '#2C2C2E'
				},
				muted: {
					DEFAULT: '#8A8A8D',
					dark: '#636366'
				}
			}
		}
	},
	plugins: []
};
