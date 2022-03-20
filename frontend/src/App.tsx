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
import { Sun, MoonStars } from "tabler-icons-react";
import StartPage from "./startPage/StartPage";

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
        <QueryClientProvider client={queryClient}>
          <AppShell
            padding="md"
            header={
              <Header height={60}>
                <Group sx={{ height: "100%" }} px={20} position="apart">
                  <Text size="xl" weight={800} color="red">
                    Raspberry Pi stock
                  </Text>

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
            <StartPage />
          </AppShell>
          <ReactQueryDevtools />
        </QueryClientProvider>
      </MantineProvider>
    </ColorSchemeProvider>
  );
}

export default App;
