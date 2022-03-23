import React, { useEffect, useState } from "react";
import {
  createStyles,
  Table,
  ScrollArea,
  UnstyledButton,
  Group,
  Text,
  Center,
  TextInput,
  Checkbox,
  ActionIcon,
  Container,
  Modal,
  Button,
  CloseButton,
} from "@mantine/core";
import {
  Selector,
  ChevronDown,
  ChevronUp,
  Search,
  ExternalLink,
  Check,
  CircleX,
  Notification,
} from "tabler-icons-react";
import { Website } from "./StartPage";
import AlertForm, { AlertType } from "./AlertForm";

const useStyles = createStyles((theme) => ({
  th: {
    padding: "0 !important",
  },

  control: {
    width: "100%",
    padding: `${theme.spacing.xs}px ${theme.spacing.md}px`,

    "&:hover": {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[6]
          : theme.colors.gray[0],
    },
  },

  icon: {
    width: 21,
    height: 21,
    borderRadius: 21,
  },

  rowSelected: {
    backgroundColor:
      theme.colorScheme === "dark"
        ? theme.fn.rgba(theme.colors[theme.primaryColor][7], 0.2)
        : theme.colors[theme.primaryColor][0],
  },
}));

interface TableSortProps {
  data: Website[];
  onModalSubmit: (
    websites: Website[],
    token: string,
    value: AlertType,
    captcha: string
  ) => void;
}

interface ThProps {
  children: React.ReactNode;
  reversed: boolean;
  sorted: boolean;
  onSort(): void;
}

function Th({ children, reversed, sorted, onSort }: ThProps) {
  const { classes } = useStyles();
  const Icon = sorted ? (reversed ? ChevronUp : ChevronDown) : Selector;
  return (
    <th className={classes.th}>
      <UnstyledButton onClick={onSort} className={classes.control}>
        <Group position="apart">
          <Text weight={500} size="sm">
            {children}
          </Text>
          <Center className={classes.icon}>
            <Icon size={14} />
          </Center>
        </Group>
      </UnstyledButton>
    </th>
  );
}

function filterData(data: Website[], search: string) {
  const keys = Object.keys(data[0]);
  const query = search.toLowerCase().trim();
  return data.filter((item) =>
    keys.some((key) => {
      const realKey = key as keyof Website;
      return String(item[realKey]).toLowerCase().includes(query);
    })
  );
}

function sortData(
  data: Website[],
  payload: { sortBy: keyof Website; reversed: boolean; search: string }
) {
  if (!payload.sortBy) {
    return filterData(data, payload.search);
  }

  return filterData(
    [...data].sort((a, b) => {
      if (payload.reversed) {
        return String(b[payload.sortBy]).localeCompare(
          String(a[payload.sortBy])
        );
      }

      return String(a[payload.sortBy]).localeCompare(String(b[payload.sortBy]));
    }),
    payload.search
  );
}

export function PiTable({ data, onModalSubmit }: TableSortProps) {
  const [showSelect, setShowSelect] = useState(false);
  const [opened, setOpened] = useState(false);
  const { classes, cx } = useStyles();
  const [selection, setSelection] = useState<String[]>([]);
  const [search, setSearch] = useState("");
  const [sortedData, setSortedData] = useState(data);
  const [sortBy, setSortBy] = useState<keyof Website>("in_stock");
  const [reverseSortDirection, setReverseSortDirection] = useState(true);

  const setSorting = (field: keyof Website) => {
    const reversed = field === sortBy ? !reverseSortDirection : false;
    setReverseSortDirection(reversed);
    setSortBy(field);
    setSortedData(sortData(data, { sortBy: field, reversed, search }));
  };

  useEffect(() => {
    setSortedData(
      sortData(data, { sortBy, reversed: reverseSortDirection, search })
    );
  }, [data, sortBy, reverseSortDirection, search]);

  const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { value } = event.currentTarget;
    setSearch(value);
    setSortedData(
      sortData(data, { sortBy, reversed: reverseSortDirection, search: value })
    );
  };
  const toggleRow = (id: string) =>
    setSelection((current) =>
      current.includes(id)
        ? current.filter((item) => item !== id)
        : [...current, id]
    );
  const toggleAll = () =>
    setSelection((current) =>
      current.length === data.length ? [] : data.map((item) => item.url)
    );

  const rows = sortedData.map((item) => {
    const selected = selection.includes(item.url);
    return (
      <tr key={item.url} className={cx({ [classes.rowSelected]: selected })}>
        {showSelect ? (
          <td>
            <Checkbox
              checked={selection.includes(item.url)}
              onChange={() => toggleRow(item.url)}
              transitionDuration={0}
            />
          </td>
        ) : (
          <></>
        )}
        <td>{item.name}</td>
        <td>{item.type}</td>
        <td>{item.ram + " Gig"}</td>
        <td>{item.shop}</td>
        <td>{item.price_string}</td>
        <td>
          {item.in_stock ? (
            <Check size={20} strokeWidth={2} color={"#40bf75"} />
          ) : (
            <CircleX size={20} strokeWidth={2} color={"red"} />
          )}
        </td>
        <td>
          <ActionIcon
            variant="default"
            onClick={() => window.open(item.url, "_blank")}
            size={30}
          >
            <ExternalLink size={16} />
          </ActionIcon>
        </td>
      </tr>
    );
  });

  return (
    <Container size="xl">
      <Modal
        centered
        opened={opened}
        onClose={() => setOpened(false)}
        title="Get Notified When a Pi is in Stock"
      >
        <AlertForm
          onSubmit={(token, value, captcha) => {
            onModalSubmit(
              data.filter((item) => selection.includes(item.url)),
              token,
              value,
              captcha
            );
            setSelection([]);
            setOpened(false);
            setShowSelect(false);
          }}
          onClose={function (): void {}}
        />
      </Modal>
      <ScrollArea>
        <Group position="apart">
          {showSelect ? (
            <Group>
              <ActionIcon
                color="green"
                variant="filled"
                disabled={selection.length === 0}
                onClick={() => setOpened(true)}
              >
                <Check />
              </ActionIcon>
              <CloseButton
                color="red"
                variant="filled"
                onClick={() => {
                  setShowSelect(false);
                  setSelection([]);
                }}
              />
            </Group>
          ) : (
            <Group>
              <Button
                leftIcon={<Notification />}
                color="green"
                onClick={() => setShowSelect(true)}
              >
                Add Notification
              </Button>
            </Group>
          )}
          <TextInput
            placeholder="Search by any field"
            icon={<Search size={14} />}
            value={search}
            onChange={handleSearchChange}
          />
        </Group>
        <Table
          horizontalSpacing="xs"
          verticalSpacing="xs"
          sx={{ minWidth: 700 }}
        >
          <thead>
            <tr>
              {showSelect ? (
                <th style={{ width: 40 }}>
                  <Checkbox
                    onChange={toggleAll}
                    checked={selection.length === data.length}
                    indeterminate={
                      selection.length > 0 && selection.length !== data.length
                    }
                    transitionDuration={0}
                  />
                </th>
              ) : (
                <></>
              )}
              <Th
                sorted={sortBy === "name"}
                reversed={reverseSortDirection}
                onSort={() => setSorting("name")}
              >
                Name
              </Th>
              <Th
                sorted={sortBy === "type"}
                reversed={reverseSortDirection}
                onSort={() => setSorting("type")}
              >
                Type
              </Th>
              <Th
                sorted={sortBy === "ram"}
                reversed={reverseSortDirection}
                onSort={() => setSorting("ram")}
              >
                RAM
              </Th>
              <Th
                sorted={sortBy === "shop"}
                reversed={reverseSortDirection}
                onSort={() => setSorting("shop")}
              >
                Shop
              </Th>
              <Th
                sorted={sortBy === "price_string"}
                reversed={reverseSortDirection}
                onSort={() => setSorting("price_string")}
              >
                Price
              </Th>
              <Th
                sorted={sortBy === "in_stock"}
                reversed={reverseSortDirection}
                onSort={() => setSorting("in_stock")}
              >
                Stock
              </Th>
              <th>Open</th>
            </tr>
          </thead>
          <tbody>
            {rows.length > 0 ? (
              rows
            ) : (
              <tr>
                <td colSpan={Object.keys(data[0]).length}>
                  <Text weight={500} align="center">
                    Nothing found
                  </Text>
                </td>
              </tr>
            )}
          </tbody>
        </Table>
      </ScrollArea>
    </Container>
  );
}
