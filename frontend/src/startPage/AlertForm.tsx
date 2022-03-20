import { useForm } from "@mantine/form";
import {
  Box,
  TextInput,
  Button,
  Group,
  Select,
  Code,
  InputWrapper,
} from "@mantine/core";
import { useState } from "react";
import { Recaptcha } from "./ReCaptcha";
export enum AlertType {
  pushover = "1",
  mail = "2",
  webhook = "3",
}

type AlertFormProps = {
  onSubmit: (token: string, value: AlertType, captcha: string) => void;
  onClose: () => void;
};

function AlertForm({ onSubmit }: AlertFormProps) {
  const [value, setValue] = useState(AlertType.pushover);
  const [captcha, setCaptcha] = useState("");

  const webhookForm = useForm({
    initialValues: { website: "" },
    validate: (values) => ({
      website: validURL(values.website) ? null : "Invalid URL",
    }),
  });
  const emailForm = useForm({
    initialValues: { email: "" },
    validate: (values) => ({
      email: /^\S+@\S+$/.test(values.email) ? null : "Invalid email",
    }),
  });

  const pushHoverForm = useForm({
    initialValues: { token: "" },
    validate: (values) => ({
      token: values.token.length < 20 ? "Too short token" : undefined,
    }),
  });

  function validURL(url: string) {
    var pattern = new RegExp(
      "^(https?:\\/\\/)?" + // protocol
        "((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|" + // domain name
        "((\\d{1,3}\\.){3}\\d{1,3}))" + // OR ip (v4) address
        "(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*" + // port and path
        "(\\?[;&a-z\\d%_.~+=-]*)?" + // query string
        "(\\#[-a-z\\d_]*)?$",
      "i"
    ); // fragment locator
    return !!pattern.test(url);
  }

  const exampleData = `
{
  "id": "0c275d95-34fd-4a9e-b38d-efea950c1818",
  "name": "Raspberry Pi 4 Computer Modell B, 2GB RAM",
  "shop": "berrybase",
  "url": "https://www.berrybase.de/raspberry-pi/raspberry-pi-computer/boards/raspberry-pi-4-computer-modell-b-2gb-ram",
  "type": "Pi 4",
  "ram": 2,
  "in_stock": false,
  "price_string": "49,80 €",
  "stock_number": 0,
  "time": "14:09:00"
}
`;
  const renderForm = (alertType: AlertType) => {
    switch (alertType) {
      case AlertType.pushover:
        return (
          <form
            onSubmit={pushHoverForm.onSubmit((values) =>
              onSubmit(values.token, value, captcha)
            )}
          >
            <TextInput
              label="Your Pushover Token"
              placeholder="Pushover Token"
              {...pushHoverForm.getInputProps("token")}
            />
            <Group position="right" mt="md">
              <Button type="submit" disabled={captcha === ""}>
                Save
              </Button>
            </Group>
          </form>
        );
      case AlertType.mail:
        return (
          <form
            onSubmit={emailForm.onSubmit((values) =>
              onSubmit(values.email, value, captcha)
            )}
          >
            <TextInput
              required
              label="Your Email Address"
              placeholder="Email Address"
              {...emailForm.getInputProps("email")}
            />
            <Group position="right" mt="md">
              <Button type="submit" disabled={captcha === ""}>
                Save
              </Button>
            </Group>
          </form>
        );
      case AlertType.webhook:
        return (
          <form
            onSubmit={webhookForm.onSubmit((values) =>
              onSubmit(values.website, value, captcha)
            )}
          >
            <TextInput
              required
              label="Webhook url"
              placeholder=""
              {...webhookForm.getInputProps("website")}
            />
            <InputWrapper label="Pushook Example Payload:">
              <br />
              <Code>{exampleData}</Code>
            </InputWrapper>
            <Group position="right" mt="md">
              <Button type="submit" disabled={captcha === ""}>
                Save
              </Button>
            </Group>
          </form>
        );
      default:
        return <></>;
    }
  };

  return (
    <Box sx={{ maxWidth: 340 }} mx="auto">
      <Select
        label="Select how you want to be notified when a pi is available:"
        placeholder="Pick one"
        onChange={(value) => setValue(value as AlertType)}
        value={value}
        data={[
          { value: AlertType.mail, label: "Mail" },
          { value: AlertType.webhook, label: "Webhook" },
          { value: AlertType.pushover, label: "Pushover" },
        ]}
      />
      {renderForm(value)}
      <Recaptcha
        action="test"
        sitekey={
          process.env.REACT_APP_SITEKEY ||
          "6Lf_0u4eAAAAACtgHlwaZZgKhZHFZb2ecit9zt1F"
        }
        onToken={(token: string) => setCaptcha(token)}
      />
    </Box>
  );
}
export default AlertForm;
