import { useMutation, useQuery } from "react-query";
import { PiTable } from "./PiTable";
import { AlertType } from "./AlertForm";

export type Website = {
  id: string;
  url: string;
  name: string;
  type: string;
  ram: string;
  shop: string;
  price_string: string;
  time: string;
  in_stock: boolean;
};

export type Recipient = {
  webhook: string;
  pushover: string;
  email: string;
};

type Task = {
  website: Website;
  destination: number;
  recipient: Recipient;
};

type AddTasks = {
  tasks: Task[];
  captcha: string;
};

const createTasks = async (data: AddTasks) => {
  const response = await fetch("/api/v1/alert", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  if (response.ok) {
    return response.json();
  }
  throw new Error("Error Creating Tasks");
};

function StartPage() {
  const { mutate } = useMutation(createTasks);
  const { isLoading, error, data } = useQuery<Website[]>(
    "status",
    () =>
      fetch("/api/v1/status").then((res) => res.json()),
    {
      refetchOnWindowFocus: true,
      refetchInterval: 1000,
      refetchOnMount: true,
      refetchOnReconnect: true,
      cacheTime: 5,
    }
  );
  if (isLoading) {
    return <div>Loading...</div>;
  }
  if (error) {
    return <div>Error: {error} </div>;
  }
  const onModalSubmit = (
    websites: Website[],
    token: string,
    value: AlertType,
    captcha: string
  ) => {
    const tasks: Task[] = websites.map((website) => ({
      website,
      destination: Number(value),
      recipient: {
        webhook: value === AlertType.webhook ? token : "",
        pushover: value === AlertType.pushover ? token : "",
        email: value === AlertType.mail ? token : "",
      },
    }));
    mutate({ tasks: tasks, captcha });
  };
  return <>{data && <PiTable data={data} onModalSubmit={onModalSubmit} />}</>;
}

export default StartPage;
