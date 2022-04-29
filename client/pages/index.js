import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import MockupSvg from '../public/svgs/mockup.js'

export default function Home() {
  return (
    <div className={styles.pageWrapper}>
      <Head>
        <title>Habitmo</title>
        <meta name="description" content="Build up your habits with the Habitmo Habit tracker" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div className={styles.pageContainer} >
        <div className={styles.headerWrapper}>
          <h1>Build up Your Habits</h1>
          <div className={styles.buttonWrapper}>
            <button>Demo</button>
            <button>Sign in / Sign up</button>
          </div>
        </div>
        <div className={styles.mockupWrapper}>
          <div className={styles.mockupContainer}>
            <MockupSvg/>
          </div>
          <div className={styles.mockupShadow}/>
        </div>
      </div>
    </div>
  )
}
