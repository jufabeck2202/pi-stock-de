import React, { useState } from "react";
import { QueryClient, QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";
import {
  ActionIcon,
  AppShell,
  ColorScheme,
  ColorSchemeProvider,
  Group,
  Header,
  MantineProvider,
  Text,
} from "@mantine/core";
import { Sun, MoonStars, BrandGithub } from "tabler-icons-react";
import StartPage from "./startPage/StartPage";
import { Logo } from "./Logo";
import { NotificationsProvider } from "@mantine/notifications";
import { Routes, Route } from "react-router-dom";
import VerifyPage from "./VerifyMail/VerifyPage";

const queryClient = new QueryClient();
function App() {
  const [colorScheme, setColorScheme] = useState<ColorScheme>("dark");
  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === "dark" ? "light" : "dark"));

  return (
    <ColorSchemeProvider
      colorScheme={colorScheme}
      toggleColorScheme={toggleColorScheme}
    >
      <MantineProvider theme={{ colorScheme }}>
        <NotificationsProvider>
          <QueryClientProvider client={queryClient}>
            <AppShell
              padding="md"
              header={
                <Header height={60}>
                  <Group sx={{ height: "100%" }} px={20} position="apart">
                    <Group>
                      <Logo />
                      <Text size="xl" weight={800} color="red">
                        Raspberry-Pi Stock
                      </Text>
                    </Group>
                    <Group>
                      <ActionIcon
                        variant="default"
                        onClick={() =>
                          window.open(
                            "https://github.com/jufabeck2202/pi-stock-de",
                            "_blank"
                          )
                        }
                        size={30}
                      >
                        <BrandGithub size={16} />
                      </ActionIcon>
                      <ActionIcon
                        variant="default"
                        onClick={() => toggleColorScheme()}
                        size={30}
                      >
                        {colorScheme === "dark" ? (
                          <Sun size={16} />
                        ) : (
                          <MoonStars size={16} />
                        )}
                      </ActionIcon>
                    </Group>
                  </Group>
                </Header>
              }
              styles={(theme) => ({
                main: {
                  backgroundColor:
                    theme.colorScheme === "dark"
                      ? theme.colors.dark[8]
                      : theme.colors.gray[0],
                },
              })}
            >
              <Routes>
                <Route path="/" element={<StartPage />} />
                <Route path="/verify/:email" element={<VerifyPage />} />
              </Routes>
            </AppShell>
            <ReactQueryDevtools />
          </QueryClientProvider>
        </NotificationsProvider>
      </MantineProvider>
    </ColorSchemeProvider>
  );
}

export default App;
