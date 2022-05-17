import Link from 'next/link'
import css from '../styles/Home.module.css'
import { Mockup } from '../public/svgs'
import { useRouter } from 'next/router'
import { DemoLogin } from '../common/services/auth'
import { localLogin } from '../common/utils'

export default function Home() {
	const router = useRouter()

  async function demoLogin() {
    await localLogin()
    await DemoLogin()
    router.push('/dashboard')
  }

  return (
    <div className={css.pageWrapper}>
      <div className={css.pageContainer} >
        <div className={css.headerWrapper}>
          <h1>Build up Your Habits</h1>
          <div className={css.buttonWrapper}>
            <a onClick={demoLogin}> Demo </a>
            <Link href="/auth" passHref> Sign in / Sign up </Link>
          </div>
        </div>
        <div className={css.mockupWrapper}>
          <div className={css.mockupContainer}>
            <Mockup/>
          </div>
          <div className={css.mockupShadow}/>
        </div>
        <div className={css.footer}>
          <h6>
            Designed and developed by: 
            <a href='https://github.com/labadcloyd' target="_blank"> Cloyd Abad</a>
          </h6>
        </div>
      </div>
    </div>
  )
}
