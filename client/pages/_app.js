import '../styles/globals.css'

function MyApp({ Component, pageProps }) {
  return (
    <div>
      <meta name="viewport" content="width=device-width, initial-scale=1.0"></meta>
      <Component {...pageProps} />
    </div>
  )

}

export default MyApp
