import type { Metadata } from "next";
import { NextIntlClientProvider } from "next-intl";
import { getMessages } from "next-intl/server";
import { Providers } from "./providers";
import "./globals.css";

export const metadata: Metadata = {
  title: "IELTS Reading bootstrap",
  description: "Deterministic Reading training bootstrap slice",
};

export default async function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  const messages = await getMessages();

  return (
    <html lang="en">
      <body>
        <NextIntlClientProvider messages={messages}>
          <Providers>{children}</Providers>
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
