//Component to display on verify
import { Text } from "@mantine/core";
import { useQuery } from "react-query";
import { useParams } from "react-router-dom";

type UnsubscribeEmailResponse = {
  error: boolean;
  msg: string;
};
export default function UnsubscribeEmailPage() {
  const { email } = useParams();
  const { isLoading, error, data } = useQuery<UnsubscribeEmailResponse>(
    "status",
    () => fetch("/api/v1/unsubscribe/" + email).then((res) => res.json()),
    {
      refetchOnWindowFocus: false,
      refetchOnMount: false,
      refetchOnReconnect: false,
      cacheTime: 0,
    }
  );
  if (isLoading) {
    return <Text color={"white"}>Loading...</Text>;
  }
  if (error) {
    return <Text color={"white"}>Error verifying email </Text>;
  }
  if (data?.error) {
    return <Text color={"white"}>Error: {data.msg} </Text>;
  }

  if (!data?.error) {
    return (
      <Text color={"white"} size="xl">
        {data?.msg}: You will no longer receive Emails
      </Text>
    );
  }
  return (
    <Text size="xl" color={"white"}>
      Wrong Email
    </Text>
  );
}
