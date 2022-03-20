import React, { FC, useEffect, useState } from "react";
import { useRecaptcha, Badge } from "react-recaptcha-hook";

export interface RecaptchaProps {
  action: string;
  sitekey: string;
  onToken: Function;
}

export const Recaptcha: FC<RecaptchaProps> = (props) => {
  const { sitekey, action, onToken } = props;
  const execute = useRecaptcha({ sitekey, hideDefaultBadge: true });
  const [token, setToken] = useState("");
  useEffect(() => {
    const getToken = async () => {
      const token = await execute(action);
      if (token) {
        setToken(token);
        onToken(token);
      }
    };
    if (token === "") {
      getToken();
    }
  }, [action, execute, onToken]);

  return <Badge />;
};
