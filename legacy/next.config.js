/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    domains: ['preview.redd.it'], // add external domains here
  },
};

module.exports = nextConfig;