import 'moment/dist/locale/de';

import LanguageDetector from 'i18next-browser-languagedetector';
import de from './locales/de/translations.json';
import en from './locales/en/translations.json';
import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import moment from 'moment';

const fallbackLng = 'en';

i18n.on('languageChanged', (lng) => {
  moment.locale(lng);
});

i18n
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng,
    debug: !import.meta.env.PROD,
    supportedLngs: ['de', 'en'],
    resources: {
      de: {
        translation: de,
      },
      en: {
        translation: en,
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
