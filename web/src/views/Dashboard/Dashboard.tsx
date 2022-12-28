import { useTranslation } from "react-i18next";
import { Outlet } from "react-router-dom";
import MobileNavigation from "../../components/MobileNavigation/MobileNavigation";
import NavigationBar from "../../components/Navigation/NavigationBar/NavigationBar";
import { INavigationItem } from "../../components/Navigation/types";
import styles from "./Dashboard.module.scss";

function Dashboard() {
  const { t } = useTranslation("menu");

  const navigationItems: INavigationItem[] = [
    { icon: <></>, title: t("channels"), to: "/" },
    { icon: <></>, title: t("recordings"), to: "/recordings" },
    { icon: <></>, title: t("settings"), to: "/settings" },
    { icon: <></>, title: t("logout"), to: "/logout" },
  ];

  return (
    <div className={`${styles.root} ${styles.desktop}`}>
      <div className={styles.mobileNavBar}>
        <MobileNavigation items={navigationItems} />
      </div>
      <div className={styles.desktopNavBar}>
        <NavigationBar items={navigationItems} />
      </div>
      <main className={styles.main}>
          <Outlet />
      </main>
    </div>
  );
}

export default Dashboard;
