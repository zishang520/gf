import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import clsx from 'clsx';

import { Redirect } from '@docusaurus/router';
import styles from './index.module.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();
  return (
    <header className={clsx('hero hero--primary', styles.heroBanner)}>
      <div className="container">
      <p><img decoding="async" loading="lazy" alt="GoFrame#auto#100px#center" src="/gfdoc-md/assets/images/logo2-dea44e9a74af1387f11f2ede68bd4434.png" width="332" height="234"/></p>
        <p>GoFrame是一款模块化、高性能、企业级的Go基础开发框架。GoFrame是一款通用性的基础开发框架，是Golang标准库的一个增强扩展级，包含通用核心的基础开发组件，优点是实战化、模块化、文档全面、模块丰富、易用性高、通用性强、面向团队。GoFrame既可用于开发完整的工程化项目，由于框架基础采用模块化解耦设计，因此也可以作为工具库使用。</p>
        <p>如果您想使用Golang开发一个业务型项目，无论是小型还是中大型项目，GoFrame是您的不二之选。如果您想开发一个Golang组件库，GoFrame提供开箱即用、丰富强大的基础组件库也能助您的工作事半功倍。如果您是团队Leader，GoFrame丰富的资料文档、详尽的代码注释、活跃的社区成员将会极大降低您的指导成本，支持团队快速接入、语言转型与能力提升。</p>
      </div>
    </header>
  );
}

export default function Home(): JSX.Element {
  const {siteConfig} = useDocusaurusContext();
  return <Redirect to={`${siteConfig.baseUrl}docs/`} />;
  
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description="Description will go into a meta tag in <head />">
      <HomepageHeader />
      <main>
        {/* <HomepageFeatures /> */}
      </main>
    </Layout>
  );
}
