import Link from 'next/link'
import styles from '../styles/Home.module.css'
import { MockupSvg } from '../public/svgs'

export default function Home() {
  return (
    <div className={styles.pageWrapper}>
      <div className={styles.pageContainer} >
        <div className={styles.headerWrapper}>
          <h1>Build up Your Habits</h1>
          <div className={styles.buttonWrapper}>
            <Link href="/" passHref> Demo </Link>
            <Link href="/auth" passHref> Sign in / Sign up </Link>
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
