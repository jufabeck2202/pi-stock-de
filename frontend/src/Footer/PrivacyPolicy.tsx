//Component to display on verify
import { createStyles, Text } from "@mantine/core";
const useStyles = createStyles((theme) => ({
  container: {
    color: theme.colorScheme !== "dark" ? "black" : "white",
    fontFamily: " ",
  },
}));
export default function PrivacyPolicy() {
  const { classes } = useStyles();
  return (
    <div className={classes.container}>
      <Text>
        <h1>Privacy Policy for Pi-Stock</h1>
        <p>
          At pi.juli.sh, accessible from https://pi.juli.sh, one of our main
          priorities is the privacy of our visitors. This Privacy Policy
          document contains types of information that is collected and recorded
          by pi.juli.sh and how we use it.
        </p>
        <p>
          If you have additional questions or require more information about our
          Privacy Policy, do not hesitate to contact us.
        </p>
        <h2>General Data Protection Regulation (GDPR)</h2>
        <p>We are a Data Controller of your information.</p>
        <p>
          Julian Beck legal basis for collecting and using the personal
          information described in this Privacy Policy depends on the Personal
          Information we collect and the specific context in which we collect
          the information:
        </p>
        <ul>
          <li>Julian Beck needs to perform a contract with you</li>
          <li>You have given Julian Beck permission to do so</li>
          <li>
            Processing your personal information is in Julian Beck legitimate
            interests
          </li>
          <li>Julian Beck needs to comply with the law</li>
          <li>
            Julian Beck will only store your e-mail address to notify you when
            an item is in stock
          </li>

          <li>Julian Beck will not send you Marketing or advertising emails</li>
          <li>
            If a user unsubscribes or delete an Email Adddress/Notifcation
            provider, than the Data will be completly deleted and not retained
            enywhere
          </li>
        </ul>
        <p>
          Julian Beck will retain your personal information (email) only for as
          long as is necessary for the purposes set out in this Privacy Policy.
          We will retain and use your information to the extent necessary to
          comply with our legal obligations, resolve disputes, and enforce our
          policies.
        </p>
        <p>
          If you are a resident of the European Economic Area (EEA), you have
          certain data protection rights. If you wish to be informed what
          Personal Information we hold about you and if you want it to be
          removed from our systems, please contact us.
        </p>
        <p>
          In certain circumstances, you have the following data protection
          rights:
        </p>
        <ul>
          <li>
            The right to access, update or to delete the information we have on
            you.
          </li>
          <li>The right of rectification.</li>
          <li>The right to object.</li>
          <li>The right of restriction.</li>
          <li>The right to data portability</li>
          <li>The right to withdraw consent</li>
        </ul>
        <h2>Log Files</h2>
        <p>
          pi.juli.sh follows a standard procedure of using log files. These
          files log visitors when they visit websites. All hosting companies do
          this and a part of hosting services' analytics. The information
          collected by log files include internet protocol (IP) addresses,
          browser type, Internet Service Provider (ISP), date and time stamp,
          referring/exit pages, and possibly the number of clicks. These are not
          linked to any information that is personally identifiable. The purpose
          of the information is for analyzing trends, administering the site,
          tracking users' movement on the website, and gathering demographic
          information.
        </p>
        <h2>Cookies and Web Beacons</h2>
        <p>
          Like any other website, pi.juli.sh uses 'cookies'. These cookies are
          only used to identify the user as a user using googles recaptcha
          service. "Google reCAPTCHA" (hereinafter "reCAPTCHA") is implemented
          on our websites. This service is provided by Google Inc., 1600
          Amphitheatre Parkway, Mountain View, CA 94043, USA ("Google"). We use
          reCHAPTA on our websites to prevent various types of data processing
          abuse attempts (e.g. phishing). reCHAPTA analyzes various information
          of the respective visitor (e.g. IP address, mouse movements of the
          user, length of stay on the website). By using reCAPTCHA, data is
          transferred to Google, which is used to determine whether the visitor
          is a human being or a (spam) bot.
          {/* used to store information including visitors' preferences, and the pages
        on the website that the visitor accessed or visited. The information is
        used to optimize the users' experience by customizing our web page
        content based on visitors' browser type and/or other information. */}
        </p>
        <h2>Consent</h2>
        <p>
          By using our website, you hereby consent to our Privacy Policy and
          agree to its terms.
        </p>
        <h2>Contact</h2>
        For more information contact me at mail@julianbeck.com
      </Text>
    </div>
  );
}
