import { useQuery } from "react-query";
import { ActionIcon, Button, Container, Table } from "@mantine/core";
import { Check, ChevronDown, CircleX, ExternalLink } from "tabler-icons-react";


function Items() {
  const { isLoading, error, data } = useQuery(
    "status",
    () =>
      fetch("http://localhost:3001/api/v1/status").then((res) => res.json()),
    {
      refetchOnWindowFocus: true,
      refetchInterval: 100,
      refetchOnMount: true,
      refetchOnReconnect: true,
    }
  );

  if (isLoading) return <h1>"Loading..."</h1>;
  if (error) return <h1> {error}</h1>;

  const rows = data.map((pi: any) => (
    <tr key={pi.id}>
      <td>{pi.name}</td>
      <td>{pi.type}</td>
      <td>{pi.ram + " ram"}</td>
      <td>{pi.shop}</td>
      <td>{pi.price_string}</td>
      <td>
        {pi.in_stock ? (
          <Check size={20} strokeWidth={2} color={"#40bf75"} />
        ) : (
          <CircleX size={20} strokeWidth={2} color={"red"} />
        )}
      </td>
      <td>{pi.time}</td>
      <td>
        <ActionIcon
          variant="default"
          onClick={() => window.open(pi.url, "_blank")}
          size={30}
        >
          <ExternalLink size={16} />
        </ActionIcon>
      </td>
    </tr>
  ));

  return (
    <Container>
      <Table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Ram</th>
            <th>Shop</th>
            <th>Price</th>
            <th>In Stock</th>
            <th>Updated</th>
            <th>Open</th>
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </Table>
    </Container>
  );
}

export default Items;
