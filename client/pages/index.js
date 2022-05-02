import Link from 'next/link'
import css from '../styles/Home.module.css'
import { Mockup } from '../public/svgs'

export default function Home() {
  return (
    <div className={css.pageWrapper}>
      <div className={css.pageContainer} >
        <div className={css.headerWrapper}>
          <h1>Build up Your Habits</h1>
          <div className={css.buttonWrapper}>
            <Link href="/" passHref> Demo </Link>
            <Link href="/auth" passHref> Sign in / Sign up </Link>
          </div>
        </div>
        <div className={css.mockupWrapper}>
          <div className={css.mockupContainer}>
            <Mockup/>
          </div>
          <div className={css.mockupShadow}/>
        </div>
      </div>
    </div>
  )
}
