import i18n from "i18next";
import { initReactI18next } from "react-i18next";

import Backend from "i18next-http-backend";
import LanguageDetector from "i18next-browser-languagedetector";

const fallbackLng = "de";

i18n
  .use(Backend)
  .use(LanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng,
    debug: true,
    defaultNS: "common",
    supportedLngs: ["de", "en"],

    interpolation: {
      escapeValue: false,
    },
  });

export default i18n;
