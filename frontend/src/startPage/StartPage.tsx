import { useMutation, useQuery } from "react-query";
import { PiTable } from "./PiTable";
import { AlertType } from "./AlertForm";
import { useNotifications } from "@mantine/notifications";
import { Check, X } from "tabler-icons-react";
import { FooterCentered } from "../footer";

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
  unix_time: number;
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

type DeleteTask = {
  destination: number;
  recipient: Recipient;
  captcha: string;
};

const createTasks = async (data: AddTasks) => {
  const response = await fetch("http://localhost:3001/api/v1/alert", {
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

const deleteTasks = async (data: DeleteTask): Promise<number> => {
  const response = await fetch("http://localhost:3001/api/v1/alert", {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });
  if (response.ok) {
    return response.json();
  }
  throw new Error("Error Deleting Tasks");
};

function StartPage(): JSX.Element {
  const { mutateAsync: addNotifcations } = useMutation(createTasks, {
    onSuccess: async (data, input) => {
      notifications.showNotification({
        title: "Successfully Subscribed",
        message: `You have successfully subscribed to ${input.tasks.length} websites`,
        color: "green",
        icon: <Check />,
      });
    },
    onError: async (error) => {
      console.log(error);
      notifications.showNotification({
        title: "Error disabeling Notifications",
        message: `${error}`,
        color: "red",
        icon: <X />,
      });
    },
  });
  const { mutateAsync: deleteNotifications } = useMutation(deleteTasks, {
    onSuccess: async (data) => {
      notifications.showNotification({
        title: "Successfully disabled all Notification",
        message: `Notifications for ${data} Item's removed `,
        color: "green",
        icon: <Check />,
      });
    },
    onError: async (error) => {
      console.log(error);
      notifications.showNotification({
        title: "Error disabeling Notifications",
        message: `${error}`,
        color: "red",
        icon: <X />,
      });
    },
  });
  const notifications = useNotifications();
  const { isLoading, error, data } = useQuery<Website[]>(
    "status",
    () =>
      fetch("http://localhost:3001/api/v1/status").then((res) => res.json()),
    {
      refetchOnWindowFocus: true,
      refetchInterval: 10000,
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

  const onModalSubmit = async (
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
    await addNotifcations({ tasks: tasks, captcha });
  };

  const onModalDelete = async (
    token: string,
    value: AlertType,
    captcha: string
  ) => {
    const deleteTask: DeleteTask = {
      destination: Number(value),
      recipient: {
        webhook: value === AlertType.webhook ? token : "",
        pushover: value === AlertType.pushover ? token : "",
        email: value === AlertType.mail ? token : "",
      },
      captcha,
    };
    await deleteNotifications(deleteTask);
  };

  return (
    <>
      {data && (
        <PiTable
          data={data}
          onModalSubmit={onModalSubmit}
          onRemoveModalSubmit={onModalDelete}
        />
      )}
    </>
  );
}

export default StartPage;
