//Component to display on verify
import { Text } from "@mantine/core";
import { useQuery } from "react-query";
import { useParams } from "react-router-dom";

type VerifyRespose = {
  error: boolean;
  msg: string;
};
export default function VerifyPage() {
  const { email } = useParams();
  const { isLoading, error, data } = useQuery<VerifyRespose>(
    "status",
    () => fetch("/api/v1/verify/" + email).then((res) => res.json()),
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
        Your Email "{data?.msg}" has been verified!
      </Text>
    );
  }
  return (
    <Text size="xl" color={"white"}>
      Wrong Email
    </Text>
  );
}
