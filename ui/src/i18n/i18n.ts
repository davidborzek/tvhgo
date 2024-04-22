import LanguageDetector from 'i18next-browser-languagedetector';
import de from './locales/de/translations.json';
import en from './locales/en/translations.json';
import es from './locales/es/translations.json';
import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import moment from 'moment/min/moment-with-locales';

const fallbackLng = 'en';

i18n.on('languageChanged', (lng) => {
  const timeLocale = localStorage.getItem('time_locale');

  if (timeLocale) {
    moment.locale(timeLocale);
  } else {
    moment.locale(lng);
  }
});

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng,
    debug: !import.meta.env.PROD,
    supportedLngs: ['de', 'en', 'es'],
    resources: {
      de: {
        translation: de,
      },
      en: {
        translation: en,
      },
      es: {
        translation: es,
      },
    },

    interpolation: {
      escapeValue: false,
    },
    react: {
      useSuspense: false,
    },
  });

i18n.services.formatter?.add('moment', (value, _lng, options) => {
  return moment(new Date(value * 1000)).format(options.format);
});

i18n.services.formatter?.add('event_duration', (value) => {
  return `${Math.floor((value.endsAt - value.startsAt) / 60)}`;
});

export default i18n;
