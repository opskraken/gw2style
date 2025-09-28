import { useEffect, useRef } from 'react';
import Head from 'next/head';
import Header from '@components/Header';
import Footer from '@components/Footer';
import Image from 'next/image';
import Link from 'next/link';
import styles from '../styles/Home.module.css';

export default function Home({ posts }) {
  const gridRef = useRef(null);
  const masonryRef = useRef(null);

  useEffect(() => {
    const initMasonry = async () => {
      if (
        typeof window !== 'undefined' &&
        gridRef.current &&
        posts &&
        posts.length > 0
      ) {
        try {
          // Dynamic imports for client-side only
          const Masonry = (await import('masonry-layout')).default;
          const imagesLoaded = (await import('imagesloaded')).default;

          // Destroy existing masonry instance
          if (masonryRef.current) {
            masonryRef.current.destroy();
          }

          // Wait for all images to load before initializing masonry
          imagesLoaded(gridRef.current, () => {
            console.log('Images loaded, initializing masonry...');
            masonryRef.current = new Masonry(gridRef.current, {
              itemSelector: `.${styles.card}`,
              columnWidth: `.${styles.gridSizer}`,
              percentPosition: true,
              transitionDuration: '0.3s',
              fitWidth: true,
            });
            console.log('Masonry initialized');
          });
        } catch (error) {
          console.log('Masonry not available, using CSS Grid fallback');
        }
      }
    };

    // Small delay to ensure DOM is ready
    const timer = setTimeout(initMasonry, 100);

    return () => {
      clearTimeout(timer);
      if (masonryRef.current) {
        masonryRef.current.destroy();
      }
    };
  }, [posts]);

  console.log('Posts data:', posts); // Debug log

  return (
    <div className="container">
      <Head>
        <title>gw2style</title>
        <link rel="icon" href="/favicon.png" />
      </Head>

      <Header />

      <div ref={gridRef} className={styles.grid}>
        <div className={styles.gridSizer}></div>
        {posts.map((post) => (
          <Link
            key={post.id}
            href={`/posts/${post.id}`}
            className={styles.card}
          >
            <div>
              <div className={styles.imageWrapper}>
                <Image
                  src={post.thumbnail}
                  alt={post.title}
                  width={300}
                  height={0}
                  style={{
                    width: '100%',
                    height: 'auto',
                    display: 'block',
                  }}
                />
              </div>
              <div className={styles.info}>
                <h2 className={styles.title}>{post.title}</h2>
                <p className={styles.views}>{post.views} views</p>
              </div>
            </div>
          </Link>
        ))}
      </div>

      <Footer />
    </div>
  );
}

export async function getServerSideProps() {
  try {
    // const res = await fetch(
    //   `http://localhost:8888/.netlify/functions/getPosts`
    // );
    const res = await fetch(
      `https://gw2style.netlify.app/.netlify/functions/getPosts`
    );
    const posts = await res.json();

    return {
      props: {
        posts: posts || [],
      },
    };
  } catch (error) {
    console.error('Error fetching posts:', error);
    return {
      props: {
        posts: [],
      },
    };
  }
}
